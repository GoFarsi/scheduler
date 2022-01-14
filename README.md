# Scheduler
[![Go Reference](https://pkg.go.dev/badge/github.com/Ja7ad/Scheduler.svg)](https://pkg.go.dev/github.com/Ja7ad/Scheduler)

[Scheduler](https://pkg.go.dev/github.com/Ja7ad/Scheduler) package is a zero-dependency scheduling library for Go

# Install
```console
go get -u github.com/Ja7ad/Scheduler
```

# Features
- Scheduling your functions
- In a scheduler instance, you can run more than one thousand jobs at a time (Max Job is `10.000`)
- In the form of Safely, you can run your jobs and if a panic occurs, your jobs will be recovered and reported to the console (`func (j *Job) DoJobSafely(jobFunction interface{}, params ...interface{}) error`)
- Multiple scheduler instances can be run simultaneously

# Example

**[More Example in](./_example)**

```go
package main

import (
	"fmt"
	"github.com/Ja7ad/Scheduler"
)

var (
	Sched = Scheduler.NewScheduler()
)

func main() {
	if err := Sched.Every(5).Second().Do(Greeting); err != nil {
		panic(err)
	}

	<-Sched.Start()
}

func Greeting() {
	fmt.Println("Hello, World!")
}
```

# Contributing

We'd love to see your contribution to the scheduler! you can contribute by following these steps :
1. Fork the repository
2. Create a new branch
3. Make your changes
4. Commit your changes
5. Push your changes to the remote repository
6. Create a pull request