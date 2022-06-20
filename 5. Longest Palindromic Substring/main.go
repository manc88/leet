//https://leetcode.com/problems/longest-palindromic-substring/
package main

func longestPalindrome(s string) string {

	slen := len(s)

	checkPoly := func(ind int) string {

		for i, j := ind, ind; ; i, j = i-1, j+1 {

			if i < 0 {
				return s[i+1 : j]
			}

			if j >= slen {
				return s[i+1 : j]
			}

			if s[i] != s[j] {
				return s[i+1 : j]
			}

		}

	}

	checkPolyEven := func(ind int) string {

		for i, j := ind-1, ind; ; i, j = i-1, j+1 {

			if i < 0 {
				return s[i+1 : j]
			}

			if j >= slen {
				return s[i+1 : j]
			}

			if s[i] != s[j] {
				return s[i+1 : j]
			}

		}

	}

	maxpolly := ""
	polly := ""
	for i := 0; i < slen; i++ {
		polly = checkPoly(i)
		if len(polly) > len(maxpolly) {
			maxpolly = polly
		}
		polly = checkPolyEven(i)
		if len(polly) > len(maxpolly) {
			maxpolly = polly
		}
	}

	return maxpolly
}
