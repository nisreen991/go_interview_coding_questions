package impl

func DeletionAtHead(head *Node) *Node {
	if head == nil {
		return head
	}
	return head.Next
}
