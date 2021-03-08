package march

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	return getIntersectionNode(headA, headB)
}

func getIntersectionNode(headA, headB *ListNode) (res *ListNode) {
	headA = invert(headA)
	for res = headB; res != nil && res.Val > 0; res = res.Next {
	}
	headB = invert(headB)
	headA = invert(headA)
	if res != nil {
		res.Val = -res.Val
	} else {
		headB = invert(headB)
	}
	return
}

func Node(val int, next *ListNode) *ListNode {
	return &ListNode{val, next}
}

func Nxt(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}

func invert(head *ListNode) (res *ListNode) {
	if head == nil {
		return
	}
	for next := head.Next; head != nil; head, next = next, Nxt(next) {
		head.Val = -head.Val
		head.Next, res = res, head
	}
	return
}
