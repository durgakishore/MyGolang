package main

import "fmt"

/*


Data types specify the type of data that a valid Go variable can hold.
In Go language, the type is divided into four categories which are as follows:

Basic type: Numbers, strings, and booleans come under this category.
Aggregate type: Array and structs come under this category.
Reference type: Pointers, slices, maps, functions, and channels come under this category.
Interface type

*/

func dataTypes() {
	fmt.Println("DataTypes Output")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%T %T %T %T\n", i, f, b, s)              // Prints type of the variable
	fmt.Printf("%v   %v      %v  %q     \n", i, f, b, s) //prints initial value of the variable
}
