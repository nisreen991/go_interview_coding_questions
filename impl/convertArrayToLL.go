package impl

type Node struct {
	Value int
	Next  *Node
}

func ConvertArrayToLL(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{Value: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &Node{Value: val}
		current = current.Next
	}
	return head
}
