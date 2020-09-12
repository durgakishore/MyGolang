package main

import "fmt"

func main() {

	name := "KISHORE"
	fmt.Println("name is :", name)

	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

}
