package impl

func Execute() {
	arr := []int{1, 2, 3, 4, 5}
	dllHead := ConvertArrayToDLL(arr)
	_ = dllHead // Use dllHead as needed
	TraversalInDLL(dllHead)
}
