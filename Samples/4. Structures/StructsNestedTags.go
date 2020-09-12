package main

import (
	"fmt"
)

type salary struct {
	Basic, HRA, TA float64
}

type employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []salary
}

func nestedStructs() {

	e := employee{}
	e.FirstName = "Kishore"
	e.LastName = "yerubandi"
	e.Email = "kishore@gmail.com"
	e.Age = 31
	e.MonthlySalary = []salary{salary{Basic: 15000.00, HRA: 5000.00, TA: 2000.00}}
	fmt.Println(e)

	var sal []salary
	sal = e.MonthlySalary
	fmt.Println("Basic: ", sal[0].Basic)
}
