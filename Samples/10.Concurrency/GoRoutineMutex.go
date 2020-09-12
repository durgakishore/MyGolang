package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	balance int
)

func goRoutineMutexExample() {

	fmt.Printf("Mutex example\n")

	done := make(chan bool)

	go deposit(1000, done)
	go withdraw(500, done)

	<-done
	<-done
	fmt.Printf("New Balance %d\n", balance)
}

func deposit(amount int, done chan bool) {

	mutex.Lock()

	fmt.Printf("Depositing %d to account with balance %d\n", amount, balance)
	balance += amount
	mutex.Unlock()
	done <- true

}

func withdraw(amount int, done chan bool) {

	time.Sleep(1000 * time.Millisecond)
	mutex.Lock()
	if balance < amount {
		fmt.Printf("Withdraw failed because of low balance\n")
	} else {
		fmt.Printf("Withdrawing %d from account with balance %d\n", amount, balance)
		balance -= amount
	}
	mutex.Unlock()
	done <- true
}
