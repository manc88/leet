package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{1, 2, 3, 4, 5}, 3))
}

func findMaxAverage(nums []int, k int) float64 {
	max := math.MinInt

	cur := 0
	for i := 0; i < len(nums); i++ {

		cur += nums[i]

		if i+1 == k {
			if cur > max {
				max = cur
			}
		}

		if i+1 > k {
			cur -= nums[i-k]
			if cur > max {
				max = cur
			}
		}

	}

	return float64(max) / float64(k)
}

func sum(q []int) int {
	res := 0
	for _, v := range q {
		res += v
	}

	return res
}
