package main

import "fmt"

func twoSum(nums []uint8, target uint8) []uint8 {

	inpLen := uint8(len(nums))
	hashmap := make(map[uint8]uint8)
	var i uint8 = 0

	for ; i < inpLen; i++ {
		if elem, ok := hashmap[target-nums[i]]; ok {
			if i != elem {
				return []uint8{i, elem}
			}
		}
		hashmap[nums[i]] = i
	}
	return nil
}

func main() {
	test := []uint8{3, 3}
	ans := uint8(6)
	result := twoSum(test, ans)
	fmt.Println(result)
}
