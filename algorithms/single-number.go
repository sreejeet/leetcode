package main

import "fmt"

func singleNumber(nums []int) int {
	num := 0
	for _, v := range nums {
		num ^= v
	}
	return num
}

func main() {
	test := []int{1, 1, 3, 2, 2}
	fmt.Println("Testing", test)
	result := singleNumber(test)
	fmt.Printf("Yields %v\n", result)
}
