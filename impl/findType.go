package impl

import "fmt"

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

func DisplayType() {
	findType("Indravadan Sarabhai")
	findType(900)
	findType(false)
}
