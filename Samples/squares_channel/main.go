package main

import "fmt"

func digits(number int, dcn chan int) {
	for number != 0 {
		digit := number % 10
		dcn <- digit
		number = number / 10
	}
	close(dcn)
}
func calcSquares(number int, squareop chan int) {

	sum := 0
	dcn := make(chan int)
	go digits(number, dcn)
	for digit := range dcn {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dcn := make(chan int)
	go digits(number, dcn)
	for digit := range dcn {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {

	number := 589
	squareop := make(chan int)
	cubeop := make(chan int)

	go calcSquares(number, squareop)
	go calcCubes(number, cubeop)

	squares, cubes := <-squareop, <-cubeop

	fmt.Println("Final output is ", squares+cubes)
}
