package main

import "fmt"

func main() {

	myMap := map[string]int{
		"Kishore": 1000,
		"Rishika": 2000,
	}

	myMap["Amma"] = 3000

	fmt.Println(myMap)
	myMap2 := make(map[int]int)

	myMap2[1] = 10
	myMap2[2] = 20
	myMap2[3] = 30
	myMap2[4] = 40
	fmt.Println(myMap2)
}
