package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {

	if singleInstance == nil {

		once.Do(
			func() {
				singleInstance = &single{}
				fmt.Println("new instance is created")
			})
	} else {
		fmt.Println("instance is already created")
	}
	return singleInstance
}
