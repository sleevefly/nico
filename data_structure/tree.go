package data_structure

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func PreorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	PreorderTraversal(node.Left)
	PreorderTraversal(node.Right)
}
