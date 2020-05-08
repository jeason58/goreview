package linkedlist

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) Print() {
	for p := l; p != nil; p = p.Next {
		fmt.Printf("%d\t", p.Val)
	}
	fmt.Println()
}


/*
a1 := []int{1, 2, 4}
a2 := []int{1, 3, 4}
*/
func InitList(a1, a2 []int) (l1, l2 *ListNode) {
	for i, p := 0, l1; i < len(a1); i++ {
		cur := &ListNode{
			Val:  a1[i],
			Next: nil,
		}

		if p == nil {
			p = cur
			l1 = p
		} else {
			p.Next = cur
			p = p.Next
		}
	}

	for i, p := 0, l2; i < len(a2); i++ {
		cur := &ListNode{
			Val:  a2[i],
			Next: nil,
		}

		if p == nil {
			p = cur
			l2 = p
		} else {
			p.Next = cur
			p = p.Next
		}
	}

	return l1, l2
}


func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h := new(ListNode)
	cur, p, q := h, l1, l2
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else if p.Val > q.Val {
			cur.Next = q
			q = q.Next
		} else {
			cur.Next = p
			p = p.Next
			cur = cur.Next
			cur.Next = q
			q = q.Next
		}

		cur = cur.Next
	}

	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}

	return h.Next
}
