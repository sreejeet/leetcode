package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, sum := nums[0], nums[0]
	for i := int16(1); i < int16(len(nums)); i++ {
		if sum+nums[i] >= nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	println("Started...")
	test := []int{-2, 3}
	fmt.Printf("Testing: %v\n", test)
	result := maxSubArray(test)
	fmt.Printf("Yields : %v\n", result)
}
