package interview

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}


func inOrderTreeByRecursive(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		inOrderTreeByRecursive(root.Left)
	}

	fmt.Println(root.Val)

	if root.Right != nil {
		inOrderTreeByRecursive(root.Right)
	}
}

func getHOfTree(root *TreeNode) (h int) {
	if root == nil {
		return 0
	}

	leftH := getHOfTree(root.Left)
	rightH := getHOfTree(root.Right)

	if leftH >= rightH {
		return leftH + 1
	} else {
		return rightH + 1
	}
}

func isCircle(head *ListNode) bool {
	if head == nil {
		return false
	}

	s, f := head, head
	for f != nil && f.Next != nil {
		if f == s || f.Next == s {
			return true
		}

		s = s.Next
		f = f.Next.Next
	}

	return false
}
