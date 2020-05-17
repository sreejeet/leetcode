package main

import (
	"fmt"
)

func majorityElement(nums []int) (m int) {
	m = nums[0]
	min, count := len(nums)/2, 1
	for _, v := range nums[1:] {
		if v == m {
			count++
			if count == min+1 {
				return
			}
		} else {
			count--
			if count == 0 {
				count++
				m = v
			}
		}
	}
	return
}

func main() {
	test := []int{1, 3, 0, 7, 0, 0, 0, 0}
	fmt.Printf("Testing %v \n", test)
	result := majorityElement(test)
	fmt.Printf("Yields %v\n", result)

	test = []int{1, 3, 3}
	fmt.Printf("Testing %v \n", test)
	result = majorityElement(test)
	fmt.Printf("Yields %v\n", result)

	test = []int{1}
	fmt.Printf("Testing %v \n", test)
	result = majorityElement(test)
	fmt.Printf("Yields %v\n", result)
}

/*
[1]
[3,2,3]
[8,8,7,7,7]
[1,3,3,4,3,3]
[1, 3, 0, 7, 0, 0, 0, 0]
*/
