package main

import (
	"fmt"
	"sort"
)

func main() {
	type Employee struct {
		name   string
		salary int
	}

	employees := []Employee{
		{"A", 9000},
		{"R", 2000},
		{"T", 60000},
		{"E", 20000},
		{"O", 5000},
		{"Ru", 12000},
		{"W", 22000},
		{"ROP", 2500},
		{"Q", 33000},
		{"N", 45000},
	}

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].salary < employees[j].salary
	})

	fmt.Println("Printing employee salaries in ascending order: ")

	for _, emp := range employees {
		fmt.Println("Name and Salary: ", emp.name, emp.salary)
	}
}
