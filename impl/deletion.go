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
