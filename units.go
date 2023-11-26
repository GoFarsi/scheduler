package scheduler

import (
	"github.com/GoFarsi/scheduler/types"
	"time"
)

// Seconds set the unit with seconds
func (j *Job) Seconds() *Job {
	return j.SetJobUnit(types.Seconds)
}

// Minutes set the unit with minutes
func (j *Job) Minutes() *Job {
	return j.SetJobUnit(types.Minutes)
}

// Hours set the unit with hours
func (j *Job) Hours() *Job {
	return j.SetJobUnit(types.Hours)
}

// Days set the unit with days
func (j *Job) Days() *Job {
	return j.SetJobUnit(types.Days)
}

// Weeks set the unit with weeks
func (j *Job) Weeks() *Job {
	return j.SetJobUnit(types.Weeks)
}

// Months set the unit with months
func (j *Job) Months() *Job {
	return j.SetJobUnit(types.Months)
}

// Second sets the unit with second
func (j *Job) Second() *Job {
	j.MustInterval(1)
	return j.Seconds()
}

// Minute sets the unit  with minute, which interval is 1
func (j *Job) Minute() *Job {
	j.MustInterval(1)
	return j.Minutes()
}

// Hour sets the unit with hour, which interval is 1
func (j *Job) Hour() *Job {
	j.MustInterval(1)
	return j.Hours()
}

// Day sets the job's unit with day, which interval is 1
func (j *Job) Day() *Job {
	j.MustInterval(1)
	return j.Days()
}

// Week sets the job's unit with week, which interval is 1
func (j *Job) Week() *Job {
	j.MustInterval(1)
	return j.Weeks()
}

// Month sets the job's unit with week, which interval is 1
func (j *Job) Month() *Job {
	j.MustInterval(1)
	return j.Weeks()
}

// Weekday start job on specific weekday
func (j *Job) Weekday(startDay time.Weekday) *Job {
	j.MustInterval(1)
	j.FirstWeekDay = startDay
	return j.Weeks()
}

// GetWeekday this returns the day of the week that the job will run
// Only use this when .Weekday(...) is called on the job.
func (j *Job) GetWeekday() time.Weekday {
	return j.FirstWeekDay
}

// Monday set the start day with monday
// - s.Every(1).Monday().Do(task)
func (j *Job) Monday() (job *Job) {
	return j.Weekday(time.Monday)
}

// Tuesday sets the job start day tuesday
func (j *Job) Tuesday() *Job {
	return j.Weekday(time.Tuesday)
}

// Wednesday sets the job start day wednesday
func (j *Job) Wednesday() *Job {
	return j.Weekday(time.Wednesday)
}

// Thursday sets the job start day thursday
func (j *Job) Thursday() *Job {
	return j.Weekday(time.Thursday)
}

// Friday sets the job start day friday
func (j *Job) Friday() *Job {
	return j.Weekday(time.Friday)
}

// Saturday sets the job start day saturday
func (j *Job) Saturday() *Job {
	return j.Weekday(time.Saturday)
}

// Sunday sets the job start day sunday
func (j *Job) Sunday() *Job {
	return j.Weekday(time.Sunday)
}
