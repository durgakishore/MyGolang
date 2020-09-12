package main

import (
	"fmt"
	"strings"
)

func stringSplit() {

	str1 := "this is, a example ,for strings split"
	str2 := "My name is Kishore"
	str3 := "string sample"

	fmt.Println("Original strings")

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

	fmt.Println("After splitting")

	str11 := strings.Split(str1, ",")
	str22 := strings.Split(str2, "")
	str33 := strings.Split(str3, "!")

	fmt.Println(str11)
	fmt.Println("length of the str11 is ", len(str11))
	fmt.Println(str22)
	fmt.Println(str33)
}
