//https://leetcode.com/problems/3sum/
package main

import (
	"fmt"
	"sort"
	"time"
)

//hinted
func ThreeSum1(nums []int) [][]int {

	start := time.Now()
	var res = make([][]int, 0)
	l := len(nums)
	sort.Ints(nums)

	for i := 0; i < l; i++ {

		n := nums[i]
		left := i + 1
		right := l - 1

		for left < right {

			sum := nums[left] + nums[right]

			switch {
			case sum+n < 0:
				left++
			case sum+n > 0:
				right--
			default:

				res = append(res, []int{n, nums[left], nums[right]})

				leftprev := left
				for left < right && nums[left] == nums[leftprev] {
					left++
				}

				rightprev := right
				for right > left && nums[right] == nums[rightprev] {
					right--
				}
			}

			for i+1 < l && nums[i+1] == nums[i] {
				i++
			}
		}
	}

	fmt.Println(time.Since(start).Seconds(), len(res))
	return res

}

//naive
func ThreeSum(nums []int) [][]int {

	start := time.Now()
	l := len(nums)
	var m = make(map[[3]int]struct{}, 20000)
	var res = make([][]int, 0)

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			for k := j; k < l; k++ {
				if i != j && j != k && k != i && nums[i]+nums[j]+nums[k] == 0 {
					c := [3]int{nums[i], nums[j], nums[k]}
					//sort.Ints(c)
					//t := (*[3]int)(c)
					if _, ok := m[c]; !ok {
						m[c] = struct{}{}
						res = append(res, c[:])
					}
				}
			}
		}

	}

	fmt.Println(time.Since(start).Seconds(), len(res))
	return res

}

func main() {

}
