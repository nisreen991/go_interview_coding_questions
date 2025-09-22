package impl

import "fmt"

func assert(i interface{}) {
	//s := i.(int) // this line throws error if in main function concrete type of s is not int

	s, ok := i.(int) // if concrete type is int, ok is true and value is in s, if not the ok is false and value of s if 0
	fmt.Println("Interface s:", s, ok)
}

func DisplayAssert() {
	s := "Nisreen Sabir"
	assert(s)
}
