package hard

// ListNode is a node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeList creates list
func MakeList(xs []int) (res *ListNode) {
	res = nil
	for i := len(xs) - 1; i >= 0; i-- {
		res = &ListNode{Val: xs[i], Next: res}
	}
	return
}

// ToSlice converts to slice
func (list *ListNode) ToSlice() (res []int) {
	res = []int{}
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return
}
