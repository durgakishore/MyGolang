package main

import (
	"fmt"
)

func multiDimentionaArrays() {

	arr1 := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	fmt.Println(arr1)
	fmt.Println("length", len(arr1))

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(arr1[i][j])
		}
	}

	for index, value := range arr1 {

		fmt.Println("value ", value, " index ", index)
		for index1, value1 := range value {
			fmt.Println("value1 ", value1, " index1 ", index1)
		}
	}
}
