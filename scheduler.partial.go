package scheduler

import (
	"github.com/Ja7ad/scheduler/helper"
	"github.com/Ja7ad/scheduler/types"
	"time"
)

// Every schedules a new periodic job running in specific interval
func Every(interval uint64) *Job {
	return defaultScheduler.Every(interval)
}

// RunPending runs all pending jobs
func RunPending() {
	defaultScheduler.RunPending()
}

// RunAll runs all jobs
func RunAll() {
	defaultScheduler.RunAll()
}

// RunAllWithDelay runs all jobs with a delay
func RunAllWithDelay(d int) {
	defaultScheduler.RunAllWithDelay(d)
}

// Start run all jobs that are scheduled to run
func Start() chan bool {
	return defaultScheduler.Start()
}

// Clear all scheduled jobs
func Clear() {
	defaultScheduler.Clear()
}

// Remove a specific job
func Remove(j interface{}) {
	defaultScheduler.Remove(j)
}

// ChangeTimeLocation change the location of a time
func ChangeTimeLocation(location *time.Location) {
	types.TimeZone = location
	defaultScheduler.ChangeLocation(location)
}

// Scheduled checks if specific job j was already added
func Scheduled(j interface{}) bool {
	for _, jb := range defaultScheduler.JobList {
		if jb.JobFunction == helper.GetFunctionName(j) {
			return true
		}
	}
	return false
}

// NextRun gets the next running time
func NextRun() (job *Job, time time.Time) {
	return defaultScheduler.NextRun()
}

// Jobs returns the list of job from the defaultScheduler
func Jobs() []*Job {
	return defaultScheduler.Jobs()
}

// SetLocker sets a locker implementation
func SetLocker(l Locker) {
	jobLocker = l
}
