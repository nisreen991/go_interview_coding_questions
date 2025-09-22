package impl

import "fmt"

type SumCalculator interface {
	Sum() int
}

type Numbers struct {
	a int
	b int
}

func (num Numbers) Sum() int {
	return num.a + num.b
}

func DisplaySum() {
	num := Numbers{
		a: 5,
		b: 10,
	}
	var calculate SumCalculator = num
	fmt.Println(calculate.Sum())
}
