//https://leetcode.com/problems/binary-search/
package main

import "fmt"

func main() {
	fmt.Println("left-", search([]int{-1, 0, 3, 5, 9, 12}, -10))
	fmt.Println("left", search([]int{-1, 0, 3, 5, 9, 12}, -1))
	fmt.Println("center", search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println("right", search([]int{-1, 0, 3, 5, 9, 12}, 12))
	fmt.Println("right+", search([]int{-1, 0, 3, 5, 9, 12}, 20))
}

func search(nums []int, target int) int {

	ln := len(nums)
	left, mid, right := 0, ln/2, ln-1

	for left != right {

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2

	}

	if nums[mid] == target {
		return mid
	}

	return -1

}
