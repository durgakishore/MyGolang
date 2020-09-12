package main

import (
	"fmt"
)

func arrays() {
	numbers := [3]int{1, 2, 3}
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", numbers[0])
	fmt.Printf("%v\n", numbers[1])

	var n [3]int
	n[0] = 4
	n[1] = 5
	fmt.Println(n)

	var i [3]int = [3]int{7, 8, 9}
	fmt.Println(i)

	k := [3]int{3, 6}
	fmt.Println(k)

	greetings := [4]string{
		"good monring",
		"good afternoon",
		"good evening",
		"good night",
	}

	fmt.Println(greetings)
	//if we don't know the size of an array
	greetings1 := [...]string{
		"good monring",
		"good afternoon",
		"good evening",
		"good night",
	}
	fmt.Println(greetings1)

	a := [3]int{1, 2, 3}
	b := [3]int{3, 2, 1}
	c := [3]int{3, 1, 2}
	d := [3]int{1, 2, 3}

	fmt.Println("a==b", a == b)
	fmt.Println("a==c", a == c)
	fmt.Println("a==d", a == d)

	m := [...]int{1, 2, 3, 4, 5}

	for index := 0; index < len(m); index++ {
		fmt.Printf("m[%d]= %d\n", index, m[index])
	}

	//range operator will provide index and value
	for index, value := range m {
		fmt.Printf("m[%d] = %d\n", index, value)
	}

	for _, value := range m {
		fmt.Printf("%d\n", value)
	}

	//multi dimensional arrays
	abc := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println(abc)

	pqr := [...][2]int{
		[...]int{3, 2},
		[...]int{4, 7},
		[...]int{8, 9},
	}
	fmt.Println(pqr)

	
}
