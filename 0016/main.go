//https://leetcode.com/problems/3sum-closest/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//start := time.Now()

	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	//fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	//fmt.Println(threeSumClosest([]int{1, 1, -1}, 2))
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))

}

func threeSumClosest(nums []int, target int) int {

	var res = make([][]int, 0)
	l := len(nums)
	sort.Ints(nums)

	for i := 0; i < l; i++ {

		n := nums[i]
		left := i + 1
		right := l - 1

		for left < right {

			//sum := nums[left] + nums[right]

			res = append(res, []int{n, nums[left], nums[right]})

			leftprev := left
			for left < right && nums[left] == nums[leftprev] {
				left++
			}

			rightprev := right
			for right > left && nums[right] == nums[rightprev] {
				right--
			}

			// for i+1 < l && nums[i+1] == nums[i] {
			// 	i++
			// }
		}
	}

	closestsum := 0

	prevdist := math.MaxInt
	fmt.Println(res)
	for _, v := range res {

		sum := v[0] + v[1] + v[2]

		dist := int(math.Abs(float64(sum - target)))
		fmt.Println(v[0], v[1], v[2], "sum=", sum, "prevdist=", prevdist, "dist=", dist, "closestsum=", closestsum)
		if dist < prevdist {
			prevdist = dist
			closestsum = sum
		}

	}

	return closestsum

}

//naive
func threeSumClosest1(nums []int, target int) int {

	//fmt.Printf("%v %d = ", nums, target)

	closestsum := 0

	prevdist := math.MaxInt

	l := len(nums)
	var m = make(map[[3]int]struct{}, 20000)

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			for k := j; k < l; k++ {
				sum := nums[i] + nums[j] + nums[k]
				c := [3]int{nums[i], nums[j], nums[k]}
				if _, ok := m[c]; ok {
					continue
				}

				if i != j && j != k && k != i {

					dist := int(math.Abs(float64(sum - target)))
					//fmt.Println(nums[i], nums[j], nums[k], "sum=", sum, "prevdist=", prevdist, "dist=", dist, "closestsum=", closestsum)
					if dist < prevdist {
						prevdist = dist
						closestsum = sum
						m[c] = struct{}{}
					}

				}

			}
		}

	}

	return closestsum

}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
