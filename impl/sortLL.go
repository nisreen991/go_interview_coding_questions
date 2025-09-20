package impl

func SortLLBruteForce(head *Node) *Node {
	cnt0, cnt1, cnt2 := 0, 0, 0
	temp := head
	for temp != nil {
		switch temp.Value {
		case 0:
			cnt0 += 1
		case 1:
			cnt1 += 1
		default:
			cnt2 += 1
		}
		temp = temp.Next
	}

	temp = head

	for temp != nil {
		if cnt0 != 0 {
			temp.Value = 0
			cnt0 -= 1
		} else if cnt1 != 0 {
			temp.Value = 1
			cnt1 -= 1
		} else {
			temp.Value = 2
			cnt2 -= 1
		}
		temp = temp.Next
	}
	return head
}
