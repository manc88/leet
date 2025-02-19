package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5}, -9))
}

func searchInsert(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {
		median := (low + high) / 2

		if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(nums) || nums[low] != target {
		return low
	}

	return low
}
