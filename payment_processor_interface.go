package main

import "fmt"

//We will write a simple program that calculates the total expense for a company based on the salaries of the employees.

type SalaryCalculator interface {
	CalculateSalary() float64
}

type Permanent struct {
	empId      int
	baseSalary float64
	pf         float64
}

type Contract struct {
	empId      int
	baseSalary float64
}

func (p Permanent) CalculateSalary() float64 {
	return p.pf + p.baseSalary
}

func (c Contract) CalculateSalary() float64 {
	return c.baseSalary
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees
*/

func totalExpense(s []SalaryCalculator) {
	expense := 0.0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Println("Total Expense of Salary:", expense)
}

func main() {
	pemp1 := Permanent{
		empId:      1,
		baseSalary: 6000,
		pf:         420,
	}

	pemp2 := Permanent{
		empId:      2,
		baseSalary: 3000.5,
		pf:         500,
	}

	con1 := Contract{
		empId:      4,
		baseSalary: 2000,
	}
	employees := []SalaryCalculator{pemp1, pemp2, con1}
	totalExpense(employees)
}
