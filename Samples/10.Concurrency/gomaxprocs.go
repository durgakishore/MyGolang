package main

import (
	"fmt"
	"runtime"
	"sync"
)

func goMaxProcsExample() {

	multiline := `This example will demonstrate that how to synchronize multiple go routines.
	For this I have used channels effectively to synchronize the behaviour of go routines.`
	fmt.Println(multiline)

	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup

	fmt.Println()

	wg.Add(3)

	p1 := make(chan int)
	p2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			for j := 1; j <= 10; j++ {
				fmt.Printf("%d", j)
				if j == 10 {
					fmt.Println()
				}
			}
		}
		p1 <- 1
		defer wg.Done()
	}()

	go func() {

		for range p1 {
			close(p1)
			for i := 0; i < 10; i++ {
				for j := 'A'; j <= 'A'+25; j++ {
					fmt.Printf("%c ", j)
					if j == 'Z' {
						fmt.Println()
					}
				}
			}
		}
		p2 <- 1
		defer wg.Done()
	}()

	go func() {
		for range p2 {
			close(p2)
			for i := 0; i < 10; i++ {
				for j := 0; j <= 12; j++ {
					fmt.Printf("%d", j)
					if j == 12 {
						fmt.Println()
					}
				}
			}
		}
		defer wg.Done()
	}()

	wg.Wait()
}
