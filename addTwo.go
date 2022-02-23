package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) {
	head := &ListNode{Val: 0}
	n1, n2, bucket, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || bucket != 0 {
		switch {
		case l1 == nil:
			n1 = 0
		default:
			n1 = l1.Val
			l1 = l1.Next
		}
		switch {
		case l2 == nil:
			n2 = 0
		default:
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + bucket) % 10}
		bucket = (n1 + n2 + bucket) / 10
		current = current.Next
	}
	return head.Next
}
