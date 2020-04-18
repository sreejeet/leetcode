package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, n := range nums {
		if elem, ok := hashmap[target-n]; ok {
			return []int{i, elem}
		}
		hashmap[n] = i
	}
	return nil
}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	ans := 17
	result := twoSum(test, ans)
	fmt.Println(result)
}
