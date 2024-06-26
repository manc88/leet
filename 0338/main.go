package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countBits2(2))
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		res[i] = strings.Count(strconv.FormatInt(int64(i), 2), "1")
	}

	return res
}

func countBits2(n int) []int {
	res := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		j := i
		cur := 0
		for j != 0 {
			cur += j & 1
			j >>= 1
		}

		res[i] = cur
	}

	return res
}
