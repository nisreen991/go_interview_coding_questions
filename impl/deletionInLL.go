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
