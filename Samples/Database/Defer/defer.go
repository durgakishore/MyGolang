package main

import (
	"fmt"
)

func deferWithNillFunc() {

	var run func() = nil

	defer run()
	fmt.Println("runs")
}

type car struct {
	model string
}

func (c car) printCarModel() {

	fmt.Println("defer: car model", c.model)
}
func deferWithoutPointers() {

	c := car{model: "Honda Amaze"}

	defer c.printCarModel()

	c.model = "Honda City"

}

func (c *car) printModel() {

	fmt.Println("defer: car model", c.model)
}
func deferWithPointers() {

	c := car{model: "Honda Amaze"}

	defer c.printModel()

	c.model = "Honda City"

}
