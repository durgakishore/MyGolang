package main

import (
	"fmt"
	"creditcard"
)

func payment() {

	credit := creditcard.CreateCreditAccount("Kishore", "1111-2222-3333-4444", 12, 2019, 123)

	fmt.Printf("Owner Name: %v\n ", credit.OwnerName())

}
