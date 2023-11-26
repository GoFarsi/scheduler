package helper

import (
	"crypto/sha256"
	"fmt"
	"github.com/GoFarsi/scheduler/errs"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// GetFunctionHashedKey create a unique key for a function
func GetFunctionHashedKey(functionName string) string {
	hash := sha256.New()
	hash.Write([]byte(functionName))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// GetFunctionName get the name of the function
func GetFunctionName(functionName interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(functionName).Pointer()).Name()
}

// CallJobFuncWithParams call a function with parameters
func CallJobFuncWithParams(jobFunc interface{}, params []interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(jobFunc)
	if len(params) != f.Type().NumIn() {
		return nil, errs.ERROR_PARAMETER
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return f.Call(in), nil
}

// TimeFormat create a standard time format for scheduler
func TimeFormat(time string) (h, m, s int, err error) {
	timeSplit := strings.Split(time, ":")
	if len(timeSplit) < 2 || len(timeSplit) > 3 {
		return 0, 0, 0, errs.ERROR_TIME_FORMAT
	}

	if h, err = strconv.Atoi(timeSplit[0]); err != nil {
		return 0, 0, 0, err
	}

	if m, err = strconv.Atoi(timeSplit[1]); err != nil {
		return 0, 0, 0, err
	}

	if len(timeSplit) == 3 {
		if s, err = strconv.Atoi(timeSplit[2]); err != nil {
			return 0, 0, 0, err
		}
	}

	if h < 0 || h > 23 || m < 0 || m > 59 || s < 0 || s > 59 {
		return 0, 0, 0, errs.ERROR_TIME_FORMAT
	}

	return h, m, s, nil
}

// NextTick returns a pointer to a time that will run at the next tick
func NextTick() *time.Time {
	now := time.Now().Add(time.Second)
	return &now
}
