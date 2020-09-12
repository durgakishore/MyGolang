package main

import (
	"fmt"
	"sync"
)

var wgRecover sync.WaitGroup

func sampleExecution(nCount int, wg *sync.WaitGroup) {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()
	defer wg.Done()

	fmt.Println("Attempting x/(x-10) where x=", nCount, "answer is", nCount/(nCount-10))

}
func sampleRecover() {
	wgRecover.Add(15)

	for i := 0; i < 15; i++ {
		go func(j int) {
			sampleExecution(j, &wgRecover)
		}(i)
	}
	wgRecover.Wait()
}
