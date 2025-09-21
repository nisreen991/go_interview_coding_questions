package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	m1 := map[int]int{1: 1, 2: 2, 3: 9}
	m2 := map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println(mapEqual(m1, m2))
}

func mapEqual(m1, m2 map[int]int) bool {
	for key, _ := range m1 {
		if m1[key] == m2[key] {
			continue
		} else {
			return false
		}
	}
	return true
}
