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

/* Adding FreeLancer type to calculate totalExpense */
type Freelancer struct {
	empId       int
	ratePerHour float64
	noOfHours   int
}

func (p Permanent) CalculateSalary() float64 {
	return p.pf + p.baseSalary
}

func (c Contract) CalculateSalary() float64 {
	return c.baseSalary
}

func (f Freelancer) CalculateSalary() float64 {
	return f.ratePerHour * float64(f.noOfHours)
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

func assert(i interface{}) {
	//s := i.(int) // this line throws error if in main function concrete type of s is not int

	s, ok := i.(int) // if concrete type is int, ok is true and value is in s, if not the ok is false and value of s if 0
	fmt.Println("Interface s:", s, ok)
}

func findType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Printf("Thy name is %s!\n", i)
	case int:
		fmt.Printf("I am number: %d\n", i)
	case bool:
		fmt.Printf("Decide who am I: %t\n", i)
	default:
		fmt.Printf("Unknown type")
	}
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

	freelancer := Freelancer{
		empId:       5,
		ratePerHour: 8.5,
		noOfHours:   10,
	}
	employees := []SalaryCalculator{pemp1, pemp2, con1, freelancer}
	totalExpense(employees)

	s := "Nisreen Sabir"
	assert(s)

	findType("Indravadan Sarabhai")
	findType(900)
	findType(false)
}
