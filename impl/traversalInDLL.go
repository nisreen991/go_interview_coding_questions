package impl

import "fmt"

func TraversalInDLL(head *Node) {
	current := head
	for current != nil {
		// Process current node (for example, print its data)
		fmt.Printf("%d-> ", current.Data)
		current = current.Next
	}
}
