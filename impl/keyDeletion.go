package impl

func DeleteAllOccurences(head *Node, key int) *Node {
	temp := head

	nextNode := &Node{}
	prevNode := &Node{}

	for temp != nil {
		if temp.Data == key {
			if temp == head {
				head = head.Next
			}
			nextNode = temp.Next
			prevNode = temp.Prev
			if nextNode != nil {
				nextNode.Prev = prevNode
			}
			if prevNode != nil {
				prevNode.Next = nextNode
			}
			temp = nextNode
		} else {
			temp = temp.Next
		}
	}
	return head
}
