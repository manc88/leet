//https://leetcode.com/problems/remove-element/
package main

func removeElement(nums []int, val int) int {

	move := 0

	for i, v := range nums {
		//fmt.Println(i)
		nums[i-move] = v
		if v == val {
			move++
		}
	}

	return len(nums) - move

}
