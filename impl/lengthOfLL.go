package impl

func LengthOfLL(head *Node) int {
	length := 0

	for node := head; node != nil; node = node.Next {
		length += 1
	}

	return length
}
