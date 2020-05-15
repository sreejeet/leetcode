package main

import (
	"fmt"
)

func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	value := 0
	for i, v := range A {
		value = (value - (value/10)*10) + v
		res[i] = value%5 == 0
		value <<= 1
	}
	return res
}

func main() {
	test := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0}
	fmt.Printf("Testing %v\n", test)
	result := prefixesDivBy5(test)
	fmt.Printf("Yields %v\n", result)
}

/*
[1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0]
[0,1,1]
[0]
[1]
[1,0]
[0,1]
[0,0,0,0,0,0,0,0]
[1,1,1,1,1,1,1,1,1]
[0,1,0,1,0,1,0,1]
[1,0,1,0,1,0,1,0]
[1,1,1,1,0,0,0,0]
[0,0,0,1,1,1]
*/
