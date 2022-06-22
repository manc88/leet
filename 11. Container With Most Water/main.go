//https://leetcode.com/problems/container-with-most-water/
package main

import (
	"fmt"
)

func main() {

	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))

}

func maxArea(height []int) int {

	//full scan stupid search
	fmt.Println(height)
	l := len(height)
	res := 0
	// dp := make([][]int, l, l)
	// for i, _ := range dp {
	// 	dp[i] = make([]int, l, l)
	// }

	// for i := 0; i < l; i++ {
	// 	for j := l - 1; j >= 0; j-- {
	// 		tmp := min(height[j], height[i]) * abs(i, j)
	// 		if tmp > max {
	// 			max = tmp
	// 		}
	// 		dp[i][j] = tmp
	// 	}

	// }

	//ps(dp)

	// for i := 0; i < l; i++ {
	// 	for j := i; j < l; j++ {
	// 		tmp := min(height[j], height[i]) * abs(i, j)
	// 		if tmp > max {
	// 			max = tmp
	// 		}
	// 	}
	// }

	fromleft := true
	for i, j := 0, l-1; i < l; {

		if i == j {
			fmt.Println(i, "=", j)
			break
		}

		tmp := min(height[j], height[i]) * abs(i, j)
		fmt.Println("i=", i, "j=", j, min(height[j], height[i]), "x", abs(i, j), "=", tmp)
		if tmp > res {

			res = tmp
		}

		fromleft = height[j] >= height[i]

		if fromleft {
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

func max(a, b int) int {
	if a > b {
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

func ps(s [][]int) {

	for j := 0; j < len(s); j++ {
		fmt.Printf("____")
	}
	fmt.Printf("\n")
	for i := 0; i < len(s); i++ {

		for j := 0; j < len(s[i]); j++ {
			fmt.Printf("%2d |", s[i][j])
		}
		fmt.Printf("\n")
		for j := 0; j < len(s[i]); j++ {
			fmt.Printf("___|")
		}

		fmt.Printf("\n")
	}
}
