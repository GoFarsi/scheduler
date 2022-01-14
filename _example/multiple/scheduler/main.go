package main

import (
	"fmt"
	"github.com/Ja7ad/Scheduler"
)

var (
	Sched1 = Scheduler.NewScheduler()
	Sched2 = Scheduler.NewScheduler()
)

func main() {
	go sched1()
	sched2()
}

func sched1() {
	if err := Sched1.Every(5).Second().Do(Greeting); err != nil {
		panic(err)
	}
	<-Sched1.Start()
}

func sched2() {
	if err := Sched2.Every(10).Second().Do(Name, "Javad"); err != nil {
		panic(err)
	}
	<-Sched2.Start()
}

func Greeting() {
	fmt.Println("Hello, World!")
}

func Name(name string) {
	fmt.Println("Hello, " + name + "!")
}
