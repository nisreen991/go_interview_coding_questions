package impl

import "fmt"

func Execute() {
	arr := []int{1, 2, 3, 4, 5}
	ll := ConvertArrayToLL(arr)
	fmt.Println("Linked List created from array:")
	// LL Traversal
	for node := ll; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println()
	lengthOfLL := LengthOfLL(ll)
	fmt.Println("Length of LL: ", lengthOfLL)

	val := 6
	fmt.Println("Is", val, "present in LL? :", IsNodePresent(ll, val))
}
