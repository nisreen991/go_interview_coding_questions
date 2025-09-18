package impl

func InsertNodeAtHead(head *Node, val int) *Node {
	temp := &Node{Value: val, Next: head}

	return temp
}

func InsertNodeAtLast(head *Node, val int) *Node {
	if head == nil {
		return &Node{Value: val}
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Value: val}
	return head
}
