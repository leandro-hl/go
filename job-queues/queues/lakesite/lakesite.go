package lakesite

import (
	"fmt"
	"time"
)

func StringChan() {
	stringchan := make(chan string, 2)
	quitchan := make(chan bool)
	done := make(chan bool)

	go func(id string) {
		for {
			select {
			case str := <-stringchan:
				fmt.Printf("received %s: %s\n", id, str)
			case ok := <-quitchan:
				fmt.Printf("Done go func %s:\n", id)
				done <- ok
				return
			}
		}
	}("1")

	go func(id string) {
		for {
			select {
			case str := <-stringchan:
				fmt.Printf("received %s: %s\n", id, str)
			case ok := <-quitchan:
				fmt.Printf("Done go func %s:\n", id)
				done <- ok
				return
			}
		}
	}("2")

	fmt.Print("Adding some messages\n")

	stringchan <- "Hey there"
	stringchan <- "Aloha"
	stringchan <- "See you!"

	quitchan <- true
	<-done

	quitchan <- true
	<-done
	fmt.Print("Done all\n")
}

func TestDispatcherV1() {
	var jobs []Job

	job1 := func() error {
		fmt.Println("Inside: Job 1 starting...")
		time.Sleep(2 * time.Second)
		fmt.Println("Inside: Job 1 Finished.")
		return nil
	}

	job2 := func() error {
		fmt.Println("Inside: Job 2 starting...")
		time.Sleep(2 * time.Second)
		fmt.Println("Inside: Job 2 Finished.")
		return nil
	}

	job3 := func() error {
		fmt.Println("Inside: Job 3 starting...")
		time.Sleep(2 * time.Second)
		fmt.Println("Inside: Job 3 Finished.")
		return nil
	}

	jobs = append(jobs, job1, job2, job3)

	d := NewDispatcher(len(jobs))

	for _, j := range jobs {
		d.AddJob(j)
	}

	d.Start(2)
	<-d.Finished
}
