package scheduler

import (
	"fmt"
	"github.com/Ja7ad/scheduler/errs"
	"github.com/Ja7ad/scheduler/helper"
	"github.com/Ja7ad/scheduler/types"
	"log"
	"reflect"
	"time"
)

var (
	jobLocker Locker
)

type Locker interface {
	Lock(string) (bool, error)
	Unlock(string) error
}

// Job information about a job
type Job struct {
	JobFunction  string
	Functions    map[string]interface{}
	FuncParams   map[string][]interface{}
	Interval     uint64
	JobUnit      types.TimeUnit
	Tags         []string
	AtTime       time.Duration
	TimeLocation *time.Location
	LastRun      time.Time
	NextRun      time.Time
	FirstWeekDay time.Weekday
	JobError     error
	LockJob      bool
}

// NewJob creates a new job
func NewJob(interval uint64) *Job {
	return &Job{
		Interval:     interval,
		TimeLocation: types.TimeZone,
		LastRun:      time.Unix(0, 0),
		NextRun:      time.Unix(0, 0),
		FirstWeekDay: time.Sunday,
		Functions:    make(map[string]interface{}),
		FuncParams:   make(map[string][]interface{}),
		Tags:         []string{},
	}
}

// Run the job and reschedule it
func (j *Job) Run() ([]reflect.Value, error) {
	if j.LockJob {
		if jobLocker == nil {
			return nil, fmt.Errorf("%v %v", errs.ERROR_TRY_LOCK_JOB, j.JobFunction)
		}

		hashedKey := helper.GetFunctionHashedKey(j.JobFunction)
		if _, err := jobLocker.Lock(hashedKey); err != nil {
			return nil, fmt.Errorf("%v %v", errs.ERROR_TRY_LOCK_JOB, j.JobFunction)
		}

		defer jobLocker.Unlock(hashedKey)
	}
	result, err := helper.CallJobFuncWithParams(j.Functions[j.JobFunction], j.FuncParams[j.JobFunction])
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Do specify the jobFunc that should be executed every time the job runs
func (j *Job) Do(jobFunction interface{}, params ...interface{}) error {
	if j.JobError != nil {
		return j.JobError
	}

	jobType := reflect.TypeOf(jobFunction)
	if jobType.Kind() != reflect.Func {
		return errs.ERROR_NOT_A_FUNCTION
	}

	funcName := helper.GetFunctionName(jobFunction)
	j.Functions[funcName] = jobFunction
	j.FuncParams[funcName] = params
	j.JobFunction = funcName

	now := time.Now().In(j.TimeLocation)
	if !j.NextRun.After(now) {
		err := j.NextJobRun()
		if err != nil {
			return err
		}
	}
	return nil
}

// DoJobSafely does the same thing as Do, except it logs unexpected panics rather than unwinding them up the chain
func (j *Job) DoJobSafely(jobFunction interface{}, params ...interface{}) error {
	recoveryFunction := func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("internal panic happen: %v", r)
			}
		}()
		_, _ = helper.CallJobFuncWithParams(jobFunction, params)
	}
	return j.Do(recoveryFunction)
}

// At schedules the job to run at the given time
//
//	s.Every(1).Day().At("20:30:01").Do(task)
//	s.Every(1).Monday().At("20:30:01").Do(task)
func (j *Job) At(t string) *Job {
	h, m, s, err := helper.TimeFormat(t)
	if err != nil {
		j.JobError = errs.ERROR_TIME_FORMAT
		return j
	}
	j.AtTime = time.Duration(h)*time.Hour + time.Duration(m)*time.Minute + time.Duration(s)*time.Second
	return j
}

// NextJobRun sets the next run time for the job
func (j *Job) NextJobRun() error {
	now := time.Now()
	if j.LastRun == time.Unix(0, 0) {
		j.LastRun = now
	}

	periodDuration, err := j.PeriodDuration()
	if err != nil {
		return err
	}

	switch j.JobUnit {
	case types.Seconds, types.Minutes, types.Hours:
		j.NextRun = j.LastRun.Add(periodDuration)
	case types.Days:
		j.NextRun = j.RoundToMidNight(j.LastRun)
		j.NextRun = j.NextRun.Add(j.AtTime)
	case types.Weeks:
		j.NextRun = j.RoundToMidNight(j.LastRun)
		dayDiff := int(j.FirstWeekDay)
		dayDiff -= int(j.NextRun.Weekday())
		if dayDiff != 0 {
			j.NextRun = j.NextRun.Add(time.Duration(dayDiff) * 24 * time.Hour)
		}
		j.NextRun = j.NextRun.Add(j.AtTime)
	}

	// next possible schedule advance
	for j.NextRun.Before(now) || j.NextRun.Before(j.LastRun) {
		j.NextRun = j.NextRun.Add(periodDuration)
	}
	return nil
}

// PeriodDuration returns the duration of the job
func (j *Job) PeriodDuration() (time.Duration, error) {
	interval := time.Duration(j.Interval)
	var periodDuration time.Duration

	switch j.JobUnit {
	case types.Seconds:
		periodDuration = interval * time.Second
	case types.Minutes:
		periodDuration = interval * time.Minute
	case types.Hours:
		periodDuration = interval * time.Hour
	case types.Days:
		periodDuration = interval * time.Hour * 24
	case types.Weeks:
		periodDuration = interval * time.Hour * 24 * 7
	case types.Months:
		periodDuration = interval * time.Hour * 24 * 30
	default:
		return 0, errs.ERROR_JOB_PREIOD
	}
	return periodDuration, nil
}
