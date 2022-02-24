package data

//TreeNode is a not of a binary tree
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

//TNode is a constructor of a TreeNode
func TNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{val, left, right}
}

//TLeaf is a constructor of a Leaf TreeNode
func TLeaf(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

//ListNode is a node
type ListNode struct {
	Val  int
	Next *ListNode
}

//List is a ListNode constructor
func List(data []int) (res *ListNode) {
	cur := &res
	for _, x := range data {
		*cur = &ListNode{Val: x}
		cur = &(*cur).Next
	}
	return
}

//ToSlice converts list to int slice
func (list *ListNode) ToSlice() (res []int) {
	for ; list != nil; list = list.Next {
		res = append(res, list.Val)
	}
	return
}
