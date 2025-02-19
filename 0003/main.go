//https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	m := make(map[byte]bool)

	maxlen := 1

	for i := 0; i < len(s)-1; i++ {
		m = make(map[byte]bool)
		m[s[i]] = true
		curlen := 1
		for j := i + 1; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {

				if curlen > maxlen {
					maxlen = curlen
				}

				break
			}

			m[s[j]] = true
			curlen++

		}

		if curlen > maxlen {
			maxlen = curlen
		}

	}

	return maxlen

}
