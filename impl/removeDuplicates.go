package impl

func RemoveDuplicates(head *Node) *Node {
	temp := head
	for temp != nil && temp.Next != nil {
		nextNode := temp.Next
		for nextNode != nil && nextNode.Data == temp.Data {
			nextNode = nextNode.Next
		}
		temp.Next = nextNode
		if nextNode != nil {
			nextNode.Prev = temp
		}
		temp = temp.Next
	}
	return head
}
