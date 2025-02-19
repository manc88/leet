//https://leetcode.com/problems/container-with-most-water/
package main

import (
	"fmt"
)

func main() {

	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))

}

func maxArea(height []int) int {

	l := len(height)
	res := 0

	for i, j := 0, l-1; i < l; {

		if i == j {
			break
		}

		tmp := min(height[j], height[i]) * abs(i, j)
		//fmt.Println("i=", i, "j=", j, min(height[j], height[i]), "x", abs(i, j), "=", tmp)
		if tmp > res {

			res = tmp
		}

		if height[j] >= height[i] {
			i++
		} else {
			j--
		}

	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}
