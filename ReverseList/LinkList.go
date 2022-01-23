package ToolKit

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	q := (*ListNode)(nil)
	for p := head; p != nil; {
		p.Next, p, q = q, p.Next, p
	}
	return q
}
