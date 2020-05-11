package main

import "fmt"

func singleNumber(nums []int) (num int) {
	for _, v := range nums {
		num ^= v
	}
	return
}

func main() {
	test := []int{1, 1, 3, 2, 2}
	fmt.Println("Testing", test)
	result := singleNumber(test)
	fmt.Printf("Yields %v\n", result)
}
