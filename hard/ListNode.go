package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(xs []int) (res *ListNode) {
	res = nil
	for i := len(xs) - 1; i >= 0; i-- {
		res = &ListNode{Val: xs[i], Next: res}
	}
	return
}

func (list *ListNode) ToSlice() (res []int) {
	res = []int{}
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return
}
