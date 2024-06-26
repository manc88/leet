// https://leetcode.com/problems/palindrome-number/
package main

import "fmt"

func main() {
	fmt.Println(1/10, 1%10)
	fmt.Println(isPalindrome2(1111))
}

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)
	for i, j := 0, len(str)-1; j >= i; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	old := x
	new := 0
	d := 0

	for x != 0 { // 121 //12 //1
		x, d = x/10, x%10 // 12,1 //1,2 //0,1

		new = new*10 + d // 1 // 12 // 121
	}

	return new == old
}
