package main

import (
	"fmt"
	"github.com/GoFarsi/scheduler"
)

func main() {
	sched := scheduler.New()

	if err := sched.Every(5).Second().Do(Greeting); err != nil {
		panic(err)
	}
	if err := sched.Every(7).Second().Do(Name, "Javad"); err != nil {
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
