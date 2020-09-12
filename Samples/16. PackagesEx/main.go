package main

import (
	add "calculation/add"
	mul "calculation/multiply"
	sub "calculation/subtract"
	"fmt"
)

func main() {

	a, b := 10, 20
	c := add.Add(a, b)
	d := mul.Multiply(a, b)
	e := sub.Subtract(a, b)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
