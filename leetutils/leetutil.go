package leetutils

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

func PS(s [][]int) {

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
