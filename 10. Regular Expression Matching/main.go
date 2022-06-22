//https://leetcode.com/problems/regular-expression-matching/
package main

import (
	"fmt"
)

func main() {

	fmt.Println(isMatch("aaaa", "a*"))

}

func isMatch(s string, p string) bool {

	//stop rec trivial
	if len(p) == 0 {
		return len(s) == 0
	}

	m := len(s) != 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || m && isMatch(s[1:], p)
	} else {
		return m && isMatch(s[1:], p[1:])
	}
}
