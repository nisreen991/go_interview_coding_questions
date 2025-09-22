package impl

import "fmt"

func (p Person) Describe() string { //interface is implemented using value receiver
	return fmt.Sprintf("Name and age of the Person is: %s and %d", p.name, p.age)
}

func (a *Address) Describe() string { // interface is implemented using pointer receiver
	return fmt.Sprintf("Person resides in state: %s in country: %s", a.state, a.country)
}

func DisplayDescribe() {
	var d1 Describer
	p1 := Person{"Ammar Sabir", 16}
	d1 = p1
	fmt.Println(d1.Describe())

	var d2 Describer
	a := Address{"Seattle", "USA"}
	d2 = &a
	fmt.Println(d2.Describe())
}
