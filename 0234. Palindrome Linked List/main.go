//https://leetcode.com/problems/palindrome-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	if head.Next == nil {
		return true
	}
	var str [100000]int
	var x int
	for head != nil {
		str[x] = head.Val
		head = head.Next
		x++
	}

	for i, j := 0, x-1; j >= i; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}
