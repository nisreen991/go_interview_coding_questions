package impl

import "fmt"

//We will write a simple program that calculates the total expense for a company based on the salaries of the employees.

func (p Permanent) CalculateSalary() float64 { // interface implemented using value  receiver
	return p.pf + p.baseSalary
}

func (p Permanent) CalculateLeavesLeft() int {
	return p.totalLeaves - p.leavesTaken
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

func DisplaySalaryAndLeaves() {
	pemp1 := Permanent{
		empId:       1,
		baseSalary:  6000,
		pf:          420,
		totalLeaves: 40,
		leavesTaken: 21,
	}

	pemp2 := Permanent{
		empId:      2,
		baseSalary: 3000.5,
		pf:         500,
	}

	pemp3 := Permanent{8, 30000, 500, 20, 12}

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

	var l LeaveCalculator = pemp1
	fmt.Printf("Leaves Remaining for Emp ID: %d is: %d\n", pemp1.empId, l.CalculateLeavesLeft())

	var empOp EmployeeOperations = pemp3
	fmt.Println(empOp.CalculateLeavesLeft(), empOp.CalculateSalary())
}
