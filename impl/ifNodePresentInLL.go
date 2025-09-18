package impl

func IsNodePresent(head *Node, val int) bool {
	for node := head; node != nil; node = node.Next {
		if node.Value == val {
			return true
		}
	}
	return false
}
