package impl

type SalaryCalculator interface {
	CalculateSalary() float64
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Describer interface {
	Describe() string
}

type Permanent struct {
	empId       int
	baseSalary  float64
	pf          float64
	totalLeaves int
	leavesTaken int
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

type Person struct {
	name string
	age  int
}

type Address struct {
	state   string
	country string
}
