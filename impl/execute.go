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
}
