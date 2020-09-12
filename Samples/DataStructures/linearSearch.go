package main

import "fmt"

/*This technique pass over the list of elements, by using the index to move from the beginning
of the list to the end. Each element is examined and if it does not match the search item,
the next item is examined. By hopping from one item to its next, the list is passed over sequentially.*/

func linearSearch(list []int, index int) bool {

	for _, key := range list {
		if key == index {
			return true
		}
	}
	return false
}

func linearSearchExample() {

	list := []int{34, 45, 54, 66, 77, 74, 89}
	fmt.Println(linearSearch(list, 77))
}
