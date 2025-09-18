package impl

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

func ConvertArrayToDLL(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{Data: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		newNode := &Node{Data: arr[i]}
		current.Next = newNode
		newNode.Prev = current
		current = newNode
	}
	return head
}
