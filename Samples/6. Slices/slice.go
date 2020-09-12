package main

import (
	"fmt"
	"sort"
)

func sliceExample() {
	a := []int{2, 9, 8, 7}
	fmt.Println(a)
	b := a[:]
	fmt.Println(b)
	c := a[:3]
	fmt.Println(c)

	sort.Ints(a)
	fmt.Println(a)
}

func multiDimentionalSlice() {

	slice := [][]string{
		{"11", "22"}, {"55", "66"}, {"33", "44"},
	}
	fmt.Println(slice)

	sliceExample()
}

func sliceSplitting() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 19, 18, 17, 16, 15, 14, 13, 12, 11, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	b := a[:5]
	c := a[5:10]
	d := a[10:15]
	//a.Split()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
