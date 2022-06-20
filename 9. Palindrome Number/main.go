//https://leetcode.com/problems/palindrome-number/
package main

import "fmt"

func isPalindrome(x int) bool {

	str := fmt.Sprint(x)
	for i, j := 0, len(str)-1; j >= i; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true

}
