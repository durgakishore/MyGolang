package main

import "fmt"

type creditAccount struct{}

func (c *creditAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 250
}

type checkingAccount struct{}

func (c *checkingAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 150
}

type hybridAcccount struct {
	creditAccount
	checkingAccount
}

func main() {

	ha := &hybridAcccount{}
	fmt.Println(ha.creditAccount.AvailableFunds())
}
