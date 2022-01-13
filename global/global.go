package global

import "time"

type TimeUnit int

// MaxJobs maximum number of jobs
const MaxJobs = 10000

const (
	Seconds TimeUnit = iota + 1
	Minutes
	Hours
	Days
	Weeks
	Months
)

var (
	TimeZone = time.Local
)
