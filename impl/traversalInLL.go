package impl

import "fmt"

func TraversalInLL(head *Node) {
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println()
}
