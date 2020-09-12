package main

import (
	"fmt"
	"sync"
)

func samplePool() {
	var numCalcsCreated int

	calcPool := &sync.Pool{

		New: func() interface{} {

			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	//seed the pool with 4 KB

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024

	var wg5 sync.WaitGroup
	wg5.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg5.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

		}()
	}

	wg5.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
