package algorithm

/*
	arr := []int{1, 4, 3, 2, 5, 2}
	h := &ListNode{Val:arr[0]}
	for p, i := h, 1; i < len(arr); i++ {
		p.Next = &ListNode{Val: arr[i]}
		p = p.Next
	}

	h = Partition(h, 3)

	for h != nil {
		fmt.Printf("%d\t", h.Val)
		h = h.Next
	}

	分隔链表，将小于x的节点全部放在大于x的节点前面
*/


type ListNode struct {
	Val int
	Next *ListNode
}


func Partition(head *ListNode, x int) *ListNode {
	lhead, hhead := new(ListNode), new(ListNode)
	lowTail, highTail := lhead, hhead

	for p := head; p != nil; {
		if p.Val < x {
			lowTail.Next = p
			p = p.Next
			lowTail = lowTail.Next
			lowTail.Next = nil
		} else {
			highTail.Next = p
			p = p.Next
			highTail = highTail.Next
			highTail.Next = nil
		}
	}

	lowTail.Next = hhead.Next
	return lhead.Next
}
