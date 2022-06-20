//https://leetcode.com/problems/roman-to-integer/
package main

func romanToInt(s string) int {

	m := make(map[rune]int, 7)

	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	res := 0
	prev := ' '
	for _, v := range s {
		res += m[v]
		if m[prev] < m[v] {
			res += -m[prev] << 1
		}
		prev = v
	}

	return res

}
