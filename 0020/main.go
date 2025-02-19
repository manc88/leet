//https://leetcode.com/problems/valid-parentheses/
package main

func isValid(s string) bool {

	stck := make([]rune, 0)
	m := make(map[rune]rune, 3)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	c := make(map[rune]rune, 3)
	c[')'] = '('
	c[']'] = '['
	c['}'] = '{'

	for _, v := range s {

		if _, exist := m[v]; exist {
			stck = append(stck, v)
			continue
		}

		//pop

		n := len(stck) - 1

		if n < 0 {
			return false
		}

		pop := stck[n]
		stck = stck[:n]

		if c[v] != pop {
			return false
		}

	}

	return len(stck) == 0

}
