package february

import "leetcode/data"

//ListNode is just a List Node
type ListNode = data.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for l := 1; ; l *= 2 {
		cur := &head
		for (*cur) != nil {
			l1, l2 := *cur, splitAt(*cur, l)
			next := splitAt(l2, l)
			if (*cur) == head && l2 == nil {
				return head
			}
			*cur, cur = merge(l1, l2)
			*cur = next
		}
	}
}

func splitAt(head *ListNode, k int) *ListNode {
	for i := 0; i < k-1 && head != nil; i++ {
		head = head.Next
	}
	if head != nil {
		head.Next, head = nil, head.Next
	}
	return head
}

func merge(l1, l2 *ListNode) (res *ListNode, last **ListNode) {
	last = &res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			*last, l1 = l1, l1.Next
		} else {
			*last, l2 = l2, l2.Next
		}
		last = &(*last).Next
	}
	if l1 != nil {
		*last = l1
	} else {
		*last = l2
	}
	for (*last) != nil {
		last = &(*last).Next
	}
	return
}
