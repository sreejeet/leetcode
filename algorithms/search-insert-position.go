package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	// Taking care of some edge cases
	// No need to process much in these cases
	if len(nums) == 0 || target < nums[0] {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	}
	// Binary search
	start, end := 0, len(nums)-1
	for {
		i := (start + end) / 2
		// println(start, i, end)
		if nums[i] == target || i == start || i == end {
			if nums[i] < target {
				return i + 1
			}
			return i
		}
		if nums[i] < target {
			start = i
		} else {
			end = i
		}
	}
}

func main() {
	test := []int{12, 14, 15, 16, 17, 19}
	fmt.Printf("Testing: %v\n", test)
	result := searchInsert(test, 11)
	fmt.Printf("Yields : %v\n", result)

	fmt.Printf("Testing: %v\n", test)
	result = searchInsert(test, 12)
	fmt.Printf("Yields : %v\n", result)

	fmt.Printf("Testing: %v\n", test)
	result = searchInsert(test, 13)
	fmt.Printf("Yields : %v\n", result)

	fmt.Printf("Testing: %v\n", test)
	result = searchInsert(test, 18)
	fmt.Printf("Yields : %v\n", result)

	fmt.Printf("Testing: %v\n", test)
	result = searchInsert(test, 20)
	fmt.Printf("Yields : %v\n", result)
}
