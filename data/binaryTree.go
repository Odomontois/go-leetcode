package data

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func TNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{val, left, right}
}
