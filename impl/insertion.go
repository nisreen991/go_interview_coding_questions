package impl

func InsertionAtHead(head *Node, data int) *Node {
	newNode := &Node{Data: data}
	if head == nil {
		return newNode
	}
	newNode.Next = head
	head.Prev = newNode
	return newNode
}

func InsertionAtTail(head *Node, data int) *Node {
	newNode := &Node{Data: data}
	if head == nil {
		return newNode
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	newNode.Prev = curr
	return head
}

func InsertionAtPosition(head *Node, data int, position int) *Node {
	if position <= 0 {
		return head
	}
	if position == 1 {
		return InsertionAtHead(head, data)
	}
	newNode := &Node{Data: data}
	curr := head
	for i := 1; i < position-1 && curr != nil; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return head
	}
	newNode.Next = curr.Next
	newNode.Prev = curr
	if curr.Next != nil {
		curr.Next.Prev = newNode
	}
	curr.Next = newNode
	return head
}
