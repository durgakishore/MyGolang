package main

import (
	"fmt"
	"sync"
)

var (
	wg5     sync.WaitGroup
	counter int

	oddNum  = make(chan int)
	evenNum = make(chan int)
)

func printEven() {

	for range evenNum {
		counter++
		printNumber()
		if counter < 10 {
			oddNum <- 0
		}
	}
	close(oddNum)
	wg5.Done()
}
func printOdd() {
	for range oddNum {
		counter++
		printNumber()
		if counter < 10 {
			evenNum <- 0
		}
	}
	wg5.Done()
}

func printNumber() {

	if counter == 10 {
		close(evenNum)
	}
	fmt.Println(counter)
}

func generateOddEvenNumbers() {
	wg5.Add(2)
	go printEven()
	go printOdd()
	evenNum <- 0
	wg5.Wait()
}
