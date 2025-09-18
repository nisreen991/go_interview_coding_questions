package impl

func InsertNodeAtHead(head *Node, val int) *Node {
	temp := &Node{Value: val, Next: head}

	return temp
}
