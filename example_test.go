package scheduler

import "fmt"

func ExampleNew() {
	greetingFunc := func() {
		fmt.Println("Hello, World!")
	}

	sched := New()

	if err := sched.Every(5).Second().Do(greetingFunc); err != nil {
		panic(err)
	}

	<-sched.Start()
}
