package reverse

import "leetcode/hard"

type ListNode = hard.ListNode

func reverseKGroup(head *ListNode, k int) (res *ListNode) {
	res = head
	pin, cur := &res, head
	for {
		for i := 0; i < k; i++ {
			if cur == nil {
				return
			}
			cur = cur.Next
		}
		acc, tail := *pin, cur
		for i := 0; i < k; i++ {
			acc.Next, acc, tail = tail, acc.Next, acc
		}
		*pin, pin = tail, &(*pin).Next
	}
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}
