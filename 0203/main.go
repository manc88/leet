package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	fake := &ListNode{
		Next: head,
	}

	cur := head
	prev := fake
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
			cur = cur.Next
			continue
		}

		prev = cur
		cur = cur.Next
	}

	return fake.Next
}
