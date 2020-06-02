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


/******** 二面：实现一个最大栈，思路：借助一个辅助栈实现 ********/

type Stack struct {
	tab        []int
	replicaTab []int
	top        int
}

func (s *Stack) Push(n int) {
	s.top++
	s.tab[s.top] = n

	if s.top > 0 {
		if s.replicaTab[s.top-1] < n {
			s.replicaTab[s.top] = n
		} else {
			s.replicaTab[s.top] = s.replicaTab[s.top-1]
		}
	} else {
		s.replicaTab[s.top] = n
	}
}

func (s *Stack) Pop() (n int) {
	if s.top < 0 {
		return -1
	}

	n = s.tab[s.top]
	s.top--

	return n
}

func (s *Stack) Max() (n int) {
	if s.top > 0 {
		return s.replicaTab[s.top]
	}

	return -1
}