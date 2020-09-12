// Simple program to demonstrate use of Buffered Channel

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func goRoutineBufferChannel() {
	rand.Seed(time.Now().Unix())

	// Create a buffered channel to manage the employee vs project load.
	projects := make(chan string, 10)

	// Launch 5 goroutines to handle the projects.
	goRoutine.Add(5)
	for i := 1; i <= 5; i++ {
		go employee4(projects, i)
	}

	for j := 1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	// Close the channel so the goroutines will quit
	close(projects)
	goRoutine.Wait()
}

func employee4(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		// Wait for project to be assigned.
		project, result := <-projects

		if result == false {
			// This means the channel is empty and closed.
			fmt.Printf("Employee : %d : Exit\n", employee)
			return
		}

		fmt.Printf("Employee : %d : Started   %s\n", employee, project)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// Display time to wait
		fmt.Println("\nTime to sleep", sleep, "ms")

		// Display project completed by employee.
		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}

}

func example() {

	naturals := make(chan int)
	squares := make(chan int)

	//counter

	go func() {

		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//squares
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//printer

	for x := range squares {
		fmt.Println(x)
	}
}
