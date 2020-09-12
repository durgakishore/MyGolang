package main

import (
	"fmt"
	"strings"
)

func stringTimming() {

	str1 := "@@Kishore Yerubandi@@"
	str2 := "!!Kishore Yerubandi     "

	fmt.Println("strings before trimming")
	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println("strings after trimming")

	str1 = strings.Trim(str1, "@")
	str2 = strings.Trim(str2, "!")

	fmt.Println("str2 length before space removing is ", len(str2))
	str2 = strings.TrimRight(str2, " ")
	fmt.Println("str2 length after space removing is ", len(str2))

	fmt.Println(str1)
	fmt.Println(str2)
}
