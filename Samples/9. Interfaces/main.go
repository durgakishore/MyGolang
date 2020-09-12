package main

import "fmt"

type vowelsFinder interface {
	findVowels() []rune
}

type myString string

func (s myString) findVowels() []rune {

	var vowels []rune

	for _, rune := range s {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' ||
			rune == 'A' || rune == 'E' || rune == 'I' || rune == 'O' || rune == 'U' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {

	name := myString("AKishore")
	var v vowelsFinder
	v = name

	fmt.Println("vowels ", string(v.findVowels()))

	emp1 := Permanent{
		basicSalary: 10000,
		pf:          1000,
		empID:       1,
	}

	emp2 := Permanent{
		basicSalary: 20000,
		pf:          2000,
		empID:       2,
	}

	emp3 := Permanent{
		basicSalary: 30000,
		pf:          3000,
		empID:       3,
	}

	employees := []SalaryCalculator{emp1, emp2, emp3}
	TotalExpense(employees)

	var s interface{} = 52
	assert(s)

	e := Employee{
		firstName:   "Kishore",
		lastName:    "Y",
		basicPay:    25000,
		pf:          3000,
		totalLeaves: 30,
		leavesTaken: 14,
	}

	var emp EmpOperations = e
	emp.DisplaySalary()
	emp.DisplayLeaves()
	fmt.Println("Done")

}

func assert(i interface{}) {

	v, ok := i.(int)
	if ok {
		fmt.Println("type is int", v)
	}
}
