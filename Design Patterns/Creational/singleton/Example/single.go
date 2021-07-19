package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
			fmt.Println("new instance is created")
		} else {
			fmt.Println("instance is already created")
		}
	} else {
		fmt.Println("instance is already created")
	}
	return singleInstance
}
