package main

import (
	"fmt"
	"strconv"
)

// This problem can be solved using DP.
// But I kept this permutation solution as this
// is what came to my mind intuitively.
~func climbStairs(n int) (result int) {
	result = 1
	for twos := 1; twos <= n/2; twos++ {
		ones := n - twos*2
		var ways float64 = 1
		for num := ones + twos; num >= 2; num-- {
			ways *= float64(num)
			if num <= ones {
				ways /= float64(num)
			}
			if num <= twos {
				ways /= float64(num)
			}
		}
		i, _ := strconv.Atoi(fmt.Sprintf("%.0f", ways))
		result += i
	}
	return
}

func main() {
	test := 20
	println("Testing", test)
	result := climbStairs(test)
	println("Yields", result)
}
