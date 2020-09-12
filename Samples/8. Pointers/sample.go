package main

import "fmt"

func zeroValuePointer() {

	b := 45
	var a *int
	if a == nil {
	 fmt.Println("pointer a in nil")
	 a = &b
	 fmt.Println("value in a is ",a)
	 fmt.Println("value in *a is ",*a)
	}
}
