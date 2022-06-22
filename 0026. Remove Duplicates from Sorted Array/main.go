//https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package main

func removeDuplicates(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	x := nums[0]
	pos := 1
	for i := 1; i < n; i++ {
		if x == nums[i] {
			continue
		}
		nums[pos] = nums[i]
		x = nums[i]
		pos++
	}

	return pos

}
