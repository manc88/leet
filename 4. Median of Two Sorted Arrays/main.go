//https://leetcode.com/problems/median-of-two-sorted-arrays/
package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	//merge sort variant
	l1 := len(nums1) - 1
	l2 := len(nums2) - 1

	if l1 < 0 && l2 < 0 {
		return 0
	}

	reslen := l1 + l2 + 2
	res := make([]int, reslen, reslen)

	j1 := 0
	j2 := 0

	for i := 0; i < reslen && l1 >= 0 && l2 >= 0; i++ {

		if l1 >= j1 && l2 >= j2 && nums1[j1] < nums2[j2] {
			res[i] = nums1[j1]
			j1++
		} else if l1 >= j1 && l2 >= j2 && nums1[j1] > nums2[j2] {
			res[i] = nums2[j2]
			j2++
		} else if l1 >= j1 {
			res[i] = nums1[j1]
			j1++
		} else if l2 >= j2 {
			res[i] = nums2[j2]
			j2++
		}

	}

	if l1 < 0 && l2 >= 0 {
		res = nums2
	}
	if l2 < 0 && l1 >= 0 {
		res = nums1
	}

	if reslen%2 == 0 {

		return float64((res[(reslen/2)-1] + res[(reslen/2)])) / 2

	}
	return float64(res[reslen/2])

}
