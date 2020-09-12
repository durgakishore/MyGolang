package main

import (
	"fmt"
)

func zeroValueOfPointer() {

	a := 25
	var ptr *int

	if ptr == nil {

		fmt.Println("ptr is null")
		ptr = &a
		fmt.Println("ptr is :", ptr)
		fmt.Println("*ptr is :", *ptr)
		fmt.Println("&ptr is :", &ptr)
	}
}

func creatingPointerWithNew() {

	ptr := new(int)

	if ptr != nil {

		fmt.Println("ptr is :", ptr)
		fmt.Println("*ptr is :", *ptr)
		fmt.Println("&ptr is :", &ptr)
	}

	a := [3]int{89, 90, 91}
	pointerArgumentToArray(&a)
	fmt.Println(a)

	sliceArgumentToArray(a[:])
	fmt.Println(a)
}

/*Do not pass a pointer to an array as a argument to a function. Use slice instead.*/
func pointerArgumentToArray(arr *[3]int) {

	(*arr)[0] = 90

	// (*arr)[0]   and arr[0] both are same
}

func sliceArgumentToArray(slice []int) {

	slice[0] = 85

}
