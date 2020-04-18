package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	rem := x
	res := 0
	for {
		res = res*10 + (rem % 10)
		rem /= 10
		if res < int(math.Pow(-2, 31)) || res > int(math.Pow(2, 31)-1) {
			// Question requires int32 overflows be returned with 0
			return 0
		}
		if rem == 0 {
			break
		}
	}
	return res
}

func main() {
	var test int = 56780
	result := reverse(test)
	fmt.Println(test)
	fmt.Println(result)

	test = -201000000
	result = reverse(test)
	fmt.Println(test)
	fmt.Println(result)

	// Question requires int32 overflows be returned with 0
	test = -10000000000001
	result = reverse(test)
	fmt.Println(test)
	fmt.Println(result)
}
