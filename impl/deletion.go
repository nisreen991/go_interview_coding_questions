package impl

func DeletionAtHead(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.Next != nil {
		head.Next.Prev = nil
	}
	head = head.Next
	return head
}

func DeletionAtTail(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	if curr.Prev != nil {
		curr.Prev.Next = nil
	}
	return head
}

func DeletionAtPosition(head *Node, position int) *Node {
	if head == nil || position <= 0 {
		return head
	}
	if position == 1 {
		return DeletionAtHead(head)
	}
	curr := head
	for i := 1; i < position && curr != nil; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return head
	}
	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	}
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	}
	return head
}
