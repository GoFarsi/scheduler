package errs

import "errors"

var (
	ERROR_TIME_FORMAT    = errors.New("errs in time format")
	ERROR_PARAMETER      = errors.New("no parameters are adapted")
	ERROR_NOT_A_FUNCTION = errors.New("in the job queue, only functions can be scheduled")
	ERROR_JOB_PREIOD     = errors.New("no job period specified")
	ERROR_NIL_REFLECTION = errors.New("reflection cannot be used with nil parameters")
	ERROR_TRY_LOCK_JOB   = errors.New("with nil locker, trying to lock job")
)
