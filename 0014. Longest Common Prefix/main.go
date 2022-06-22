//https://leetcode.com/problems/longest-common-prefix/
package main

import "math"

func longestCommonPrefix(strs []string) string {

	res := ""
	for i, word := range strs {

		if i == 0 {
			res = word
			continue
		}

		if len(res) == 0 {
			return ""
		}

		border := int(math.Min(float64(len(res)-1), float64(len(word)-1)))
		res = res[:border+1]
		for j := border; j >= 0; j-- {
			if res[j] != word[j] {
				res = res[:j]
			}
		}
	}

	return res
}
