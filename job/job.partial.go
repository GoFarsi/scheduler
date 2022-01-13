package job

import (
	"fmt"
	"github.com/Ja7ad/Scheduler/global"
	"time"
)

// JobIsRunNow runs the job now
func (j *Job) JobIsRunNow() bool {
	return time.Now().Unix() >= j.NextRun.Unix()
}

// Tag adds a labels for a job
func (j *Job) Tag(t string, others ...string) {
	j.Tags = append(j.Tags, t)
	j.Tags = append(j.Tags, others...)
}

// UnTag removes a tag from a job
func (j *Job) UnTag(t string) {
	var newTags []string
	for _, tag := range j.Tags {
		if t != tag {
			newTags = append(newTags, tag)
		}
	}
}

// TagList returns the tags of a job
func (j *Job) TagList() []string {
	return j.Tags
}

// NextScheduledTime returns the next scheduled time of a job
func (j *Job) NextScheduledTime() time.Time {
	return j.NextRun
}

// From plans the next run of the job
func (j *Job) From(t *time.Time) *Job {
	j.NextRun = *t
	return j
}

// SetJobUnit sets the JobUnit
func (j *Job) SetJobUnit(unit global.TimeUnit) *Job {
	j.JobUnit = unit
	return j
}

// GetAtTime returns the time of day that the job will run
// s.Every(1).Day().At("20:30").GetAtTime() == "20:30"
func (j *Job) GetAtTime() string {
	return fmt.Sprintf("%d:%d", j.AtTime/time.Hour, (j.AtTime%time.Hour)/time.Minute)
}

// Location it set the location at which to interpret "At"
// s.Every(1).Day().At("20:30").Loc(time.UTC).Do(task)
func (j *Job) Location(location *time.Location) *Job {
	j.TimeLocation = location
	return j
}

// RoundToMidNight it set the job to run at the beginning of the next day
func (j *Job) RoundToMidNight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, j.TimeLocation)
}

// Err check for errors when creating the job to make sure none occurred
func (j *Job) Err() error {
	return j.JobError
}

// Lock prevents the job from running multi instances
func (j *Job) Lock() *Job {
	j.LockJob = true
	return j
}

// MustInterval set the job's unit with seconds,minutes,hours...
func (j *Job) MustInterval(i uint64) error {
	if j.Interval != i {
		return fmt.Errorf("interval must be %d", i)
	}
	return nil
}
