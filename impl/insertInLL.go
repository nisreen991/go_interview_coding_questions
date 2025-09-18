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

func InsertNodeAtPosition(head *Node, val int, pos int) *Node {
	//position is from 0 to length-1
	if pos == 0 {
		return InsertNodeAtHead(head, val)
	}
	current := head
	for i := 0; i < pos-1 && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		return head // Position is out of bounds
	}
	newNode := &Node{Value: val, Next: current.Next}
	current.Next = newNode
	return head
}
