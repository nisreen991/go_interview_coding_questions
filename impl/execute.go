package impl

import "fmt"

func Execute() {
	arr := []int{1, 2, 3, 4, 5}
	ll := ConvertArrayToLL(arr)
	fmt.Println("Linked List created from array:")

	TraversalInLL(ll)

	fmt.Println("Length of LL: ", LengthOfLL(ll))

	val := 6
	fmt.Println("Is", val, "present in LL? :", IsNodePresent(ll, val))

	var newLL *Node
	newLL = InsertNodeAtHead(ll, val)
	fmt.Println("Insertion at head of LL:")
	TraversalInLL(newLL)

	newLL = InsertNodeAtLast(newLL, val)
	fmt.Println("Insertion at end of LL:")
	TraversalInLL(newLL)

	pos := 3
	newLL = InsertNodeAtPosition(newLL, 10, pos)
	fmt.Println("Insertion at position", pos, "of LL:")
	TraversalInLL(newLL)

	newLL = DeletionAtHead(newLL)
	fmt.Println("Deletion at head of LL:")
	TraversalInLL(newLL)

	newLL = DeletionAtLast(newLL)
	fmt.Println("Deletion at end of LL:")
	TraversalInLL(newLL)

	pos = 2
	newLL = DeletionAtPosition(newLL, pos)
	fmt.Println("Deletion at position", pos, "of LL:")
	TraversalInLL(newLL)

	sortArr := []int{1, 0, 1, 2, 0, 2, 1}
	sortLLHead := ConvertArrayToLL(sortArr)
	sortLLHead = SortLLBruteForce(sortLLHead)
	TraversalInLL(sortLLHead)

}
