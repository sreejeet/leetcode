package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := m-1, n-1
	for end := len(nums1) - 1; end >= 0 && i2 >= 0; end-- {
		if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[end] = nums1[i1]
			i1--
		} else {
			nums1[end] = nums2[i2]
			i2--
		}
	}
}

func main() {
	test1 := []int{1, 3, 5, 7, 0, 0, 0, 0}
	test2 := []int{2, 4, 6, 8}
	println("Testing")
	fmt.Println(test1)
	fmt.Println(test2)
	merge(test1, 4, test2, len(test2))
	println("Yields")
	fmt.Println(test1)
	fmt.Println(test2)

	test1 = []int{1, 2, 4, 5, 0}
	test2 = []int{3}
	println("Testing")
	fmt.Println(test1)
	fmt.Println(test2)
	merge(test1, 4, test2, len(test2))
	println("Yields")
	fmt.Println(test1)
	fmt.Println(test2)

	test1 = []int{2, 0}
	test2 = []int{1}
	println("Testing")
	fmt.Println(test1)
	fmt.Println(test2)
	merge(test1, 1, test2, len(test2))
	println("Yields")
	fmt.Println(test1)
	fmt.Println(test2)

	test1 = []int{1, 2, 3, 0, 0, 0}
	test2 = []int{2, 5, 6}
	println("Testing")
	fmt.Println(test1)
	fmt.Println(test2)
	merge(test1, 3, test2, len(test2))
	println("Yields")
	fmt.Println(test1)
	fmt.Println(test2)

}

/*

[2,0]
1
[1]
1
[2,2,2,0]
3
[1]
1
[2,2,2,0]
3
[2]
1
[2,2,2,0]
3
[3]
1
[2,4,6,0,0]
3
[3,5]
2
[2,2,2,0,0]
3
[1,1]
2

*/
