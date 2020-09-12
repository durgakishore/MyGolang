package main

import "fmt"

//"reflect"

type rectangle struct {
	length   int
	breath   int
	color    string
	geometry struct {
		area      int
		perimeter int
	}
	square
}

type square struct {
	area int
}

type rectangle1 struct {
	length int
	breath int
	color  string
}

func moreStructExamples() {
	var rect rectangle
	rect.length = 10
	rect.breath = 20
	rect.color = "red"
	rect.geometry.area = rect.length * rect.breath
	rect.geometry.perimeter = 2 * (rect.length + rect.breath)
	rect.square.area = 200

	fmt.Println(rect)
	fmt.Println("area ", rect.geometry.area)
	fmt.Println("Perimater ", rect.geometry.perimeter)
	fmt.Println("sqaure area ", rect.square.area)

	var rect1 = rectangle{}
	rect1.length = 40
	rect1.breath = 30
	fmt.Println(rect1)

	var rect2 = new(rectangle1)
	rect2.length = 50
	rect2.breath = 40
	fmt.Println(rect2)

}
