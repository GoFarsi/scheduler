package main

import (
	"fmt"
	"github.com/Ja7ad/scheduler"
)

func main() {
	sched := scheduler.New()

	if err := sched.Every(5).Second().Do(Greeting); err != nil {
		panic(err)
	}

	<-sched.Start()
}

func Greeting() {
	fmt.Println("Hello, World!")
}
