package Scheduler

import (
	"github.com/Ja7ad/Scheduler/global"
	"github.com/Ja7ad/Scheduler/helper"
	"github.com/Ja7ad/Scheduler/job"
	"sort"
	"time"
)

var (
	defaultScheduler = NewScheduler()
)

// Scheduler is the main struct of the scheduler
type Scheduler struct {
	JobList  [global.MaxJobs]*job.Job // Jobs is the array of jobs
	JobSize  int                      // JobSize is the number of jobs
	Location *time.Location           // Location is the location of the scheduler
}

// NewScheduler creates a new scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{
		JobList:  [global.MaxJobs]*job.Job{},
		JobSize:  0,
		Location: global.TimeZone,
	}
}

// Jobs returns list of jobs from scheduler
func (s *Scheduler) Jobs() []*job.Job {
	return s.JobList[:s.JobSize]
}

// Len returns the number of jobs in the scheduler
func (s *Scheduler) Len() int {
	return s.JobSize
}

// Swap swaps the two jobs
func (s *Scheduler) Swap(i, j int) {
	s.JobList[i], s.JobList[j] = s.JobList[j], s.JobList[i]
}

// Less returns true if the job at index i is less than the job at index j
func (s *Scheduler) Less(i, j int) bool {
	return s.JobList[j].NextRun.Unix() >= s.JobList[i].NextRun.Unix()
}

// ChangeLocation changes the default time location of the scheduler
func (s *Scheduler) ChangeLocation(newLocation *time.Location) {
	s.Location = newLocation
}

// GetRunnableJobs returns a list of jobs that are ready to run
func (s *Scheduler) GetRunnableJobs() (runningJobs [global.MaxJobs]*job.Job, n int) {
	runnableJobs := [global.MaxJobs]*job.Job{}
	n = 0
	sort.Sort(s)
	for i := 0; i < s.JobSize; i++ {
		if s.JobList[i].JobIsRunNow() {
			runnableJobs[n] = s.JobList[i]
			n++
		} else {
			break
		}
	}
	return runnableJobs, n
}

// NextRun returns the next run time of the job at index i
func (s *Scheduler) NextRun() (*job.Job, time.Time) {
	if s.JobSize <= 0 {
		return nil, time.Now()
	}
	sort.Sort(s)
	return s.JobList[0], s.JobList[0].NextRun
}

// Every schedules a new job to run every duration
func (s *Scheduler) Every(interval uint64) *job.Job {
	j := job.NewJob(interval).Location(s.Location)
	s.JobList[s.JobSize] = j
	s.JobSize++
	return j
}

// RunPending runs all pending jobs
func (s *Scheduler) RunPending() {
	runnableJobs, n := s.GetRunnableJobs()

	if n != 0 {
		for i := 0; i < n; i++ {
			go runnableJobs[i].Run()
			runnableJobs[i].LastRun = time.Now()
			runnableJobs[i].NextJobRun()
		}
	}
}

// RunAll runs all jobs
func (s *Scheduler) RunAll() {
	s.RunAllWithDelay(0)
}

// RunAllWithDelay runs all jobs with a delay
func (s *Scheduler) RunAllWithDelay(d int) {
	for i := 0; i < s.JobSize; i++ {
		go s.JobList[i].Run()
		if 0 != d {
			time.Sleep(time.Duration(d))
		}
	}
}

// Remove job by function
func (s *Scheduler) Remove(j interface{}) {
	s.RemoveByCondition(func(someJob *job.Job) bool {
		return someJob.JobFunction == helper.GetFunctionName(j)
	})
}

// RemoveByRef removes specific job j by reference
func (s *Scheduler) RemoveByRef(j *job.Job) {
	s.RemoveByCondition(func(someJob *job.Job) bool {
		return someJob == j
	})
}

// RemoveByTag removes specific job j by tag
func (s *Scheduler) RemoveByTag(t string) {
	s.RemoveByCondition(func(someJob *job.Job) bool {
		for _, a := range someJob.Tags {
			if a == t {
				return true
			}
		}
		return false
	})
}

// RemoveByCondition removes specific job j by condition
func (s *Scheduler) RemoveByCondition(remove func(*job.Job) bool) {
	i := 0

	// Delete jobs until no more match the criteria
	for {
		found := false
		for ; i < s.JobSize; i++ {
			if remove(s.JobList[i]) {
				found = true
				break
			}
		}
		if !found {
			return
		}

		for j := i + 1; j < s.JobSize; j++ {
			s.JobList[i] = s.JobList[j]
			i++
		}
		s.JobSize--
		s.JobList[s.JobSize] = nil
	}

}

// Scheduled returns true if job j is scheduled
func (s *Scheduler) Scheduled(j interface{}) bool {
	for _, jb := range s.JobList {
		if jb.JobFunction == helper.GetFunctionName(j) {
			return true
		}
	}
	return false
}

// Clear removes all scheduled jobs
func (s *Scheduler) Clear() {
	for i := 0; i < s.JobSize; i++ {
		s.JobList[i] = nil
	}
	s.JobSize = 0
}

// Start starts all pending jobs
func (s *Scheduler) Start() chan bool {
	stopped := make(chan bool, 1)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				s.RunPending()
			case <-stopped:
				ticker.Stop()
				return
			}
		}
	}()

	return stopped
}
