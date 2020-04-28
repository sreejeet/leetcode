package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	var l, r int = 0, numsLen - 1
	for l <= r {
		if nums[l] != val {
			l++
		} else {
			// fmt.Printf("Stop l at %v\n", l)
			for r >= l {
				if nums[r] == val {
					r--
				} else {
					// fmt.Printf("Stop r at %v and exchanging\n", r)
					nums[l], nums[r] = nums[r], nums[l]
					l++
					r--
					break
				}
			}
		}
	}
	return l
}

func main() {
	test := []int{1, 2, 2, 2, 3, 3, 4, 5, 5, 5, 5, 5, 5, 6}
	fmt.Printf("Testing: %v\n", test)
	result := removeElement(test, 3)
	fmt.Printf("Yields : %v\n", result)
	fmt.Printf(": %v\n", test)

	test = []int{1, 2, 2, 3}
	fmt.Printf("Testing: %v\n", test)
	result = removeElement(test, 3)
	fmt.Printf("Yields : %v\n", result)
	fmt.Printf(": %v\n", test)

	test = []int{1}
	fmt.Printf("Testing: %v\n", test)
	result = removeElement(test, 3)
	fmt.Printf("Yields : %v\n", result)
	fmt.Printf(": %v\n", test)

	test = []int{3}
	fmt.Printf("Testing: %v\n", test)
	result = removeElement(test, 3)
	fmt.Printf("Yields : %v\n", result)
	fmt.Printf(": %v\n", test)

}
