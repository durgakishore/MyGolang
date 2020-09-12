package main

import (
	"fmt"
	"sync"
)

func evenNumbersTillEight(even chan int) {
	i := 2
	for i < 9 {
		even <- i
		i = i + 2
	}
	close(even)
}

func oddNumberTillEight(odd chan int) {

	i := 1
	for i < 9 {
		odd <- i
		i = i + 2
	}
	close(odd)
}

func goRoutineOddEvenSample() {

	even := make(chan int)
	odd := make(chan int)
	go evenNumbersTillEight(even)
	go oddNumberTillEight(odd)
	for {
		even, ok1 := <-even
		odd, ok2 := <-odd
		if ok1 == false && ok2 == false {
			break
		}
		fmt.Println("Received ", even, odd)
	}

	sample11()
}

func sample11() {

	c := make(chan int)
	cc := make(chan int)

	var wg sync.WaitGroup

	p := func(c chan int) {

		for v := range c {
			fmt.Println(v)
		}
		wg.Done()
	}

	wg.Add(2)

	go p(c)
	go p(cc)

	c <- 1
	c <- 5
	c <- 3
	c <- 4
	cc <- 9193

	close(c)
	close(cc)
	wg.Wait()

}
