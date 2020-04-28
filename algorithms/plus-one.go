package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	for i := len(digits) - 1; ; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
		if i == 0 {
			digits = append([]int{1}, digits...)
			return digits
		}
	}
}

func main() {
	test := []int{0, 0}
	fmt.Printf("Testing: %v\n", test)
	result := plusOne(test)
	fmt.Printf("Yields : %v\n", result)
}
