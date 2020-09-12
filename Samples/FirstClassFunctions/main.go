package main

import (
	pkg "firstclass/AnonymousFunctions"
	"fmt"
)

type add func(a int, b int) int

func addition(a int, b int) int {
	return a + b
}

func main() {

	pkg.AnonymousFunctions()

	var a add = func(a int, b int) int {
		return a + b
	}
	b := a(5, 6)
	fmt.Println("b is", b)

	c := addition

	fmt.Println(c(11, 12))

	s1 := pkg.Student{
		FirstName: "Kishore",
		LastName:  "Yerubandi",
		Grade:     "A",
		Country:   "India",
	}

	s2 := pkg.Student{
		FirstName: "Rishika",
		LastName:  "Yerubandi",
		Grade:     "B",
		Country:   "India",
	}

	s := []pkg.Student{s1, s2}

	f := pkg.Filter(s, func(s pkg.Student) bool {
		if s.Grade == "B" {
			return true
		}
		return false
	})

	fmt.Println(f)
}
