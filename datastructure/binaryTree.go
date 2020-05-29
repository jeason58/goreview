package datastructure

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func CreateTree(nums []int) *Node {
	root := &Node{Val:nums[0]}

	for i := 1; i < len(nums); i++ {
		insertTreeNode(root, nums[i])
	}

	return root
}

func insertTreeNode(root *Node, n int) {
	if root.Val == n {
		node := &Node{Val:n}
		node.Left = root.Left
		root.Left = node
	} else if root.Val > n {
		if root.Left == nil {
			root.Left = &Node{Val:n}
		} else {
			insertTreeNode(root.Left, n)
		}
	} else {
		if root.Right == nil {
			root.Right = &Node{Val:n}
		} else {
			insertTreeNode(root.Right, n)
		}
	}
}

func InOrderByRecursive(root *Node) {
	stack := make([]*Node, 0)
	top := -1

	for root != nil || top > -1 {
		for root != nil {
			top++
			if len(stack) == top {
				stack = append(stack, root)
			} else {
				stack[top] = root
			}
			root = root.Left
		}

		if top > -1 {
			root = stack[top]
			top--
			fmt.Println(root.Val)
			root = root.Right
		}
	}
}

