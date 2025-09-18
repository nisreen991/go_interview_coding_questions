package impl

import "fmt"

func Execute() {
	arr := []int{1, 2, 3, 4, 5}
	dllHead := ConvertArrayToDLL(arr)
	TraversalInDLL(dllHead)

	dllHead = DeletionAtHead(dllHead)
	fmt.Println("After deletion at head:")
	TraversalInDLL(dllHead)
}
