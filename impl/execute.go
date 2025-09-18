package impl

import "fmt"

func Execute() {
	arr := []int{1, 2, 3, 4, 5}
	dllHead := ConvertArrayToDLL(arr)
	TraversalInDLL(dllHead)

	dllHead = DeletionAtHead(dllHead)
	fmt.Println("After deletion at head:")
	TraversalInDLL(dllHead)

	dllHead = DeletionAtTail(dllHead)
	fmt.Println("After deletion at tail:")
	TraversalInDLL(dllHead)

	dllHead = DeletionAtPosition(dllHead, 2)
	fmt.Println("After deletion at position 2:")
	TraversalInDLL(dllHead)

	dllHead = InsertionAtHead(dllHead, 0)
	fmt.Println("After insertion at head:")
	TraversalInDLL(dllHead)

	dllHead = InsertionAtTail(dllHead, 6)
	fmt.Println("After insertion at tail:")
	TraversalInDLL(dllHead)

	dllHead = InsertionAtPosition(dllHead, 99, 3)
	fmt.Println("After insertion at position 3:")
	TraversalInDLL(dllHead)

	dllHead = ReverseDLL(dllHead)
	fmt.Println("After reversing the DLL:")
	TraversalInDLL(dllHead)
}
