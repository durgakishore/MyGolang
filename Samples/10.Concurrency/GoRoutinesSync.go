package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func goRoutinesSyncExample() {

	start := time.Now()

	wg.Add(4)

	firstChannel := make(chan string)
	secondChannel := make(chan string)
	thirdChannel := make(chan string)
	fourthChannel := make(chan string)

	go firstFunction(firstChannel)
	go secondFunction(secondChannel)
	go thirdFunction(thirdChannel)
	go fourthFunction(fourthChannel)

	fmt.Println(<-firstChannel)
	fmt.Println(<-secondChannel)
	fmt.Println(<-thirdChannel)
	fmt.Println(<-fourthChannel)

	wg.Wait()

	fmt.Printf("Total time to finish : %s \n", time.Since(start).String())
}

func firstFunction(ch chan string) {

	fmt.Println("Executing first function")
	time.Sleep(7 * time.Millisecond)
	defer wg.Done()
	ch <- "Execution of first functin ended"
}

func secondFunction(ch chan string) {

	fmt.Println("Executing second function")
	time.Sleep(7 * time.Millisecond)
	defer wg.Done()
	ch <- "Execution of second function ended"

}

func thirdFunction(ch chan string) {

	fmt.Println("Executing third function")
	time.Sleep(7 * time.Millisecond)
	defer wg.Done()
	ch <- "Execution of third function ended"
}

func fourthFunction(ch chan string) {

	fmt.Println("Executing fourth function")
	time.Sleep(7 * time.Millisecond)
	defer wg.Done()
	ch <- "Execution of fourth function ended"
}
