package main

import (
	"fmt"
)

// Animal struct is to be exported...
type Animal struct {
	Name   string
	Origin string
}

type animal struct {
	Name   string
	Origin string
}

type person struct {
	FirstName, LastName string
	Age                 int
}

func printStruct() {
	A := Animal{Name: "EMU", Origin: "AUSTRALIA"}
	fmt.Println(A)

	a := animal{Name: "Emu", Origin: "Austraila"}
	fmt.Println(a)

	var p person
	fmt.Println(p)

	p1 := person{FirstName: "kishore",
		LastName: "yerubandi",
		Age:      32,
	}
	ps := &p1

	fmt.Println("First Name: ", ps.FirstName)
	fmt.Println("Last Name: ", ps.LastName)
	fmt.Println("Age: ", ps.Age)
	ps.Age = 31
	fmt.Println(p1)

	ps.Age = 0
	ps.personInfo()
	ps.pointerToStructure()
	fmt.Println(p1)

}

func (p *person) personInfo() {

	if p.FirstName == "" {
		p.FirstName = "Kishna"
	}

	if p.Age == 0 {
		p.Age = 25
	}
}
func (p *person) pointerToStructure() {

	fmt.Println("First Name: ", p.FirstName)
	fmt.Println("Last Name: ", p.LastName)
	fmt.Println("Age: ", p.Age)
	p.Age = 32

}
