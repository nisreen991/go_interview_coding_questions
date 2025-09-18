package impl

func DeletionAtHead(head *Node) *Node {
	if head == nil {
		return head
	}
	return head.Next
}

func DeletionAtLast(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	// in order to delete the last node, you need to stop at second last node and point next of second last node to null
	for prev.Next.Next != nil {
		prev = prev.Next
	}
	prev.Next = nil
	return head
}

func DeletionAtPosition(head *Node, pos int) *Node {
	if pos > LengthOfLL(head) {
		return nil
	}

	if pos == 0 {
		return DeletionAtHead(head)
	}

	current := head
	for i := 0; i < pos-1 && current != nil; i++ {
		current = current.Next
	}
	if current == nil || current.Next == nil {
		return head // Position is out of bounds
	}
	current.Next = current.Next.Next
	return head
}
