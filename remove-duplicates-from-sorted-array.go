package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var i, j int = 0, 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}

func main() {
	test := []int{1, 2, 2, 2, 3, 3, 4, 5, 5, 5, 5, 5, 5, 6}
	fmt.Printf("Testing: %v\n", test)
	result := removeDuplicates(test)
	fmt.Printf("Yields : %v\n", result)
	fmt.Printf(": %v\n", test)

	test = []int{1, 3, 3, 3, 3, 3, 6}
	fmt.Printf("Testing: %v\n", test)
	result = removeDuplicates(test)
	fmt.Printf("Yields : %v\n", result)
	fmt.Printf(": %v\n", test)
}
