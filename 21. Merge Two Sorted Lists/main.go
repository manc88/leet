//https://leetcode.com/problems/merge-two-sorted-lists/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}

	if l1 != nil && l2 != nil && l1.Val >= l2.Val {

		head = &ListNode{l2.Val, nil}
		l2 = l2.Next
	} else if l1 != nil && l2 != nil && l2.Val > l1.Val {

		head = &ListNode{l1.Val, nil}
		l1 = l1.Next
	} else if l1 == nil && l2 != nil {

		head = &ListNode{l2.Val, nil}
		l2 = l2.Next
	} else if l2 == nil && l1 != nil {

		head = &ListNode{l1.Val, nil}
		l1 = l1.Next
	} else if l2 == nil && l1 == nil {
		return nil
	}

	tail := head

	for l1 != nil || l2 != nil {

		switch {

		case l1 != nil && l2 != nil && l1.Val > l2.Val:

			tail.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		case l1 != nil && l2 != nil && l1.Val == l2.Val:

			tail.Next = &ListNode{l2.Val, nil}
			tail.Next.Next = &ListNode{l1.Val, nil}
			l2 = l2.Next
			l1 = l1.Next
			tail = tail.Next
		case l1 != nil && l2 != nil && l1.Val < l2.Val:

			tail.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		case l1 == nil:

			tail.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		case l2 == nil:

			tail.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		}

		if tail != nil {
			tail = tail.Next
		}

	}

	return head

}
