package impl

func FindTailOfDll(head *Node) *Node {
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	return tail
}
