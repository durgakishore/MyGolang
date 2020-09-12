package main

import (
	"fmt"
	"sync"
)

func cashier(cashierID int, orderChannel <-chan int, wg *sync.WaitGroup) {

	for ordersProcessed := 0; ordersProcessed < 10; ordersProcessed++ {
		orderNumber := <-orderChannel
		fmt.Println("Cashier ", cashierID, "Processing order", orderNumber, "Orders Processed", ordersProcessed)
		wg.Done()
	}

}

func sampleChannel() {

	var wgOrder sync.WaitGroup
	wgOrder.Add(30)

	orderChannel := make(chan int)

	for i := 0; i < 3; i++ {
		func(i int) {
			go cashier(i, orderChannel, &wgOrder)
		}(i)
	}

	for i := 0; i < 30; i++ {

		orderChannel <- i
	}

	wgOrder.Wait()
}

func sampleChannel1() {

	f1 := func(i int) {
		fmt.Println("f1: ", i)
	}
	f2 := func(i int) {
		fmt.Println("f2: ", i*2)
	}
	f3 := func(i int) {
		fmt.Println("f3: ", i*3)
	}

	ch := make(chan func(int))
	done := make(chan bool)

	go func() {

		for fcn := range ch {
			fcn(10)
		}
		fmt.Println("exiting")
		done <- true
	}()

	//sending functions to channels

	ch <- f1
	ch <- f2
	ch <- f3

	close(ch)

	<-done
}
