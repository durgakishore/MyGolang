package main

import "fmt"

func find(num int, nums ...int) {

	fmt.Printf("Type of nums is %T\n", nums)
	bFound := false
	for i, v := range nums {
		if num == v {
			bFound = true
			fmt.Println(num, " found at index ", i)
			break
		}
	}
	if !bFound {
		fmt.Println(num, " not found in ", nums)
	}
}

func main() {

	nums := []int{10, 20, 30, 40}

	find(20, nums...)
}
