package main

import "fmt"

var mp = map[int]int{0: 0, 1: 1}

func fib(N int) int {
	for {
		if res, ok := mp[N]; ok {
			return res
		}
		mp[N] = fib(N-1) + fib(N-2)
	}
}

func main() {
	test := 30
	println("Testing", test)
	result := fib(test)
	fmt.Printf("Yields %v\n", result)
}

/*
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
*/
