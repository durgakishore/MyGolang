package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
}
func main() {
	defer r()

	a := []int{3, 4, 5}
	fmt.Println(a[3])
	fmt.Println("main execution is done")

}
