package main

import (
	"fmt"
)

// This passes all cases in 0ms, but its not the best approach
// Will update onece I figure out a better solution.
func sortColors(nums []int) {
	bound := len(nums)
	r, w, b := 0, 0, 0
	for i := 0; i < bound; i++ {
		if nums[i] == 0 {
			r++
		} else if nums[i] == 1 {
			w++
		} else {
			b++
		}
	}
	for i := 0; i < bound; i++ {
		if r != 0 {
			nums[i] = 0
			r--
		} else if w != 0 {
			nums[i] = 1
			w--
		} else if b != 0 {
			nums[i] = 2
			b--
		}
	}

}

func main() {
	test := []int{0, 1, 2, 1}
	fmt.Printf("Testing %v \n", test)
	sortColors(test)
	fmt.Printf("Yields %v\n", test)
}

/*
[0]
[1]
[2]
[0,0]
[1,1]
[2,2]
[0,1]
[1,0]
[2,0,2,1,1,0]
[2,2,1,1,0,0]
*/
