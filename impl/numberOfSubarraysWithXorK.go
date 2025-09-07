package impl

func numberofSubarraysWithXorK(arr []int, k int) int {
	xr, count := 0, 0
	preXor := make(map[int]int)
	preXor[xr] += 1 // {0,1}

	for i := 0; i < len(arr); i++ {
		xr = xr ^ arr[i]
		x := xr ^ k
		count += preXor[x]
		preXor[xr] += 1
	}
	return count
}
