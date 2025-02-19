//https://leetcode.com/problems/add-two-numbers/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	inmind := 0

	nodeRes := new(ListNode)
	node := nodeRes
	v1, v2 := 0, 0
	for ptr1, ptr2 := l1, l2; ; {

		v1 = 0
		v2 = 0
		if ptr1 != nil {
			v1 = ptr1.Val
			ptr1 = ptr1.Next
		}
		if ptr2 != nil {
			v2 = ptr2.Val
			ptr2 = ptr2.Next
		}

		sum := v1 + v2 + inmind

		if sum > 9 {
			node.Val = sum - 10
			inmind = 1
		} else {
			node.Val = sum
			inmind = 0
		}

		if ptr1 == nil && ptr2 == nil {
			if inmind == 1 {
				node.Next = new(ListNode)
				node = node.Next
				node.Val = 1
			}
			break
		}

		node.Next = new(ListNode)
		node = node.Next

	}

	return nodeRes

}
