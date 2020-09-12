package main

import "fmt"

//SalaryCalculator interface
type SalaryCalculator interface {
	CalculateSalary() int
}

//Permanent struct
type Permanent struct {
	empID       int
	basicSalary int
	pf          int
}

//Contract struct
type Contract struct {
	empID       int
	basicSalary int
}

//CalculateSalary METHOD
func (p Permanent) CalculateSalary() int {

	return p.basicSalary + p.pf
}

//CalculateSalary Method
func (c Contract) CalculateSalary() int {
	return c.basicSalary
}

//TotalExpense FUN
func TotalExpense(s []SalaryCalculator) {

	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
	}

	fmt.Println("Total salary is ", expense)
}
