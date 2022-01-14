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
	if err := Sched.Every(7).Second().Do(Name, "Javad"); err != nil {
		panic(err)
	}
	<-Sched.Start()
}

func Greeting() {
	fmt.Println("Hello, World!")
}

func Name(name string) {
	fmt.Println("Hello, " + name + "!")
}
