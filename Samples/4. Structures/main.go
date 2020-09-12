package main

import "fmt"

type emp struct {
	firstname, lastname string
	age                 int
}

type anonymous struct {
	string
	int
}

func main() {
	//structTags()
	e := emp{
		firstname: "Kishore",
		lastname:  "Y",
		age:       32,
	}
	updateStruct1(&e)
	fmt.Println(e.age)
	createAnonymousStruct()

	anon := anonymous{
		string: "hhhh",
		int:    32,
	}

	fmt.Println(anon)
}

func updateStruct1(e *emp) {

	e.age = 45
}

func createAnonymousStruct() {
	emp3 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "Rishika",
		lastname:  "Y",
		age:       4,
	}
	fmt.Println(emp3)
}
