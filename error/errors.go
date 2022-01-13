package error

import "errors"

var (
	ERROR_TIME_FORMAT    = errors.New("error in time format")
	ERROR_PARAMETER      = errors.New("no parameters are adapted")
	ERROR_NOT_A_FUNCTION = errors.New("in the job queue, only functions can be scheduled")
	ERROR_JOB_PREIOD     = errors.New("no job period specified")
	ERROR_NIL_REFLECTION = errors.New("reflection cannot be used with nil parameters")
)
