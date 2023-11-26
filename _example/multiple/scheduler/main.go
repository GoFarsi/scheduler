package main

import (
	"fmt"
	"github.com/Ja7ad/scheduler"
)

func main() {
	go sched1()
	sched2()
}

func sched1() {
	sched := scheduler.New()

	if err := sched.Every(5).Second().Do(Greeting); err != nil {
		panic(err)
	}
	<-sched.Start()
}

func sched2() {
	sched := scheduler.New()

	if err := sched.Every(10).Second().Do(Name, "Javad"); err != nil {
		panic(err)
	}
	<-sched.Start()
}

func Greeting() {
	fmt.Println("Hello, World!")
}

func Name(name string) {
	fmt.Println("Hello, " + name + "!")
}
