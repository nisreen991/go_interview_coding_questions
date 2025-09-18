package impl

func ReverseDLL(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	prev := &Node{}
	curr := head

	for curr != nil {
		prev = curr.Prev
		curr.Prev = curr.Next
		curr.Next = prev

		curr = curr.Prev
	}
	return prev.Prev
}
