package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	l1, l2 := len(nums1), len(nums2)
	mid := ((l1 + l2) / 2) + 1
	v1, v2 := 0, 0
	for i, j := 0, 0; i+j < mid; {
		if (i < l1 && j < l2 && nums1[i] <= nums2[j]) || j >= l2 {
			v1, v2 = v2, nums1[i]
			i++
		} else {
			v1, v2 = v2, nums2[j]
			j++
		}
	}
	// println(mid, mid%2, v1, v2)
	if (l1+l2)%2 == 0 {
		return float64(v1+v2) / 2.0
	}
	return float64(v2)
}

func main() {
	test1 := []int{100001}
	test2 := []int{100000}
	println("Testing")
	fmt.Println(test1)
	fmt.Println(test2)
	result := findMedianSortedArrays(test1, test2)
	println("Yields", result)
}

/*
[-5]
[-4]
[-4]
[-5]
[100000]
[100001]
[100001]
[100000]
[1,6,7,4]
[2,3,5,8]
[1,2,3,4]
[5,6,7,8]
[1,2,3,4]
[5,6,7,8]
[1,2,3,4,9]
[]
[1,9]
[2,3,4,5]
[2,3]
[1,4,5]
[2,3]
[1,4,5]
[1,2]
[3,4]
[3]
[-2,-1]
[1,3]
[2]
[1,3,5,7,9]
[2,4,6,8,10]
[]
[1]
[]
[1,2]
[]
[1,2,3]
[1,2,2,2,2,4,5,6,7,8,9,9,9]
[1,1,1,1,4,4,5,6,6,6,7,7,8,8,9,9]
*/
