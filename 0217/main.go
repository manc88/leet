// https://leetcode.com/problems/contains-duplicate/description/
package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, i := range nums {
		if _, found := m[i]; found {
			return true
		}

		m[i] = struct{}{}
	}

	return false
}
