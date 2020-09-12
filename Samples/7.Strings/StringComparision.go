package main

import (
	"bytes"
	"fmt"
	"strings"
)

func stringsComparisionTypes() {

	str1 := "kishore"
	str2 := "yerubandi"
	str3 := "Kishore"
	str4 := "Yerubandi"

	result1 := str1 == str2
	result2 := str1 == str3
	result3 := str2 == str4
	result4 := str2 == str2

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	//compare method

	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Compare(str1, str3))
	fmt.Println(strings.Compare(str2, str4))
	fmt.Println(strings.Compare(str2, str2))
}

func concatenateStrings() {

	var b bytes.Buffer

	b.WriteString("K")
	b.WriteString("i")
	b.WriteString("s")
	b.WriteString("h")
	b.WriteString("o")
	b.WriteString("r")
	b.WriteString("e")

	fmt.Println(b)          //{[75 105 115 104 111 114 101] 0 0}
	fmt.Println(b.String()) //Kishore

	str1 := "kishore"
	str2 := "yerubandi"

	result := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(result)

	slice := []string{"kishore", "yerubandi", "from", "Amalapuram"}

	fmt.Println(strings.Join(slice, " "))

}
