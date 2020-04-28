package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 10 {
		if x < 0 {
			return false
		}
		return true
	}
	if x%10 == 0 {
		return false
	}
	rev := 0
	rem := x
	for rem > 0 {
		rev = rev*10 + (rem % 10)
		rem /= 10
		if rem < rev {
			print(rem, " ", rev, "\n")
			break
		}
	}
	return rev == x
}

func main() {
	var test int = 1093901
	result := isPalindrome(test)
	fmt.Println(test)
	fmt.Println(result, "\n")

	test = 0
	result = isPalindrome(test)
	fmt.Println(test)
	fmt.Println(result, "\n")

	test = 1919191919
	result = isPalindrome(test)
	fmt.Println(test)
	fmt.Println(result, "\n")
}
