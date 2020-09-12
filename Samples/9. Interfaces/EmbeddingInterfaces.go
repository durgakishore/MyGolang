package main

import "fmt"

//SalaryCalculate interface
type SalaryCalculate interface {
	DisplaySalary()
}

//LeavesCalculate interface
type LeavesCalculate interface {
	DisplayLeaves()
}

//EmpOperations interface
type EmpOperations interface {
	SalaryCalculate
	LeavesCalculate
}

//Employee struct
type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

//DisplaySalary  func
func (emp Employee) DisplaySalary() {

	fmt.Printf("emp salary is %v", emp.basicPay+emp.pf)
}

//DisplayLeaves  func
func (emp Employee) DisplayLeaves() {

	fmt.Printf("emp remaining leaves  are %v", emp.totalLeaves-emp.leavesTaken)
}
