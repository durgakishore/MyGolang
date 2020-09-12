package main

import (
	. "fmt"
	"sync"
)

func sampleOnce() {

	var count int

	increment := func() {

		count++
	}

	decrement := func() {

		count--
	}

	var once sync.Once
	var increments sync.WaitGroup

	increments.Add(100)

	for i := 0; i < 10; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment) //It will execute only once even we run the loop
			once.Do(decrement) // It will execute this function as already one time DO execution is done already
		}()
	}

	increments.Wait()
	Printf("Count is %d", count)
}
