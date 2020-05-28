package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	nonZero := 0
	limit := len(nums)
	for i := 0; i < limit; i++ {
		for nonZero < limit && nums[nonZero] == 0 {
			nonZero++
		}
		if nonZero >= limit || i >= limit {
			break
		}
		nums[i], nums[nonZero] = nums[nonZero], nums[i]
		nonZero++
	}
}

func main() {
	test := []int{0, 1, 0, 3, 12}
	fmt.Printf("Testing %v \n", test)
	moveZeroes(test)
	fmt.Printf("Yields %v\n", test)
}

/*
[]
[1]
[0]
[0,1,0,3,12]
[0,0,0,0,0]
[1,1,1,1,1,1]
[0,1,0,1,0]
[1,0,1,0,1]
[0,0,1,1,1]
[1,1,0,1,0]
*/
