package main

import (
	"encoding/json"
	"fmt"
)

type employee1 struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func structTags() {
	fmt.Println("")

	e1 := new(employee1)

	jsonString := `{
		"FirstName":"kishore",
		"LastName" : "yerubandi",
		"city" : "Hyderabad"
	}`
	json.Unmarshal([]byte(jsonString), e1)
	fmt.Println(e1)

	e2 := new(employee1)
	e2.FirstName = "Kiran"
	e2.LastName = "abc"
	e2.City = "RJY"

	jsonStr, _ := json.Marshal(e2)
	fmt.Printf("%s\n", jsonStr)
}
