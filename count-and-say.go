package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var memo = map[int]string{
	1: "1",
}

func countAndSay(n int) string {
	if res, ok := memo[n]; ok {
		return res
	}
	var res bytes.Buffer
	count := 0
	inp := countAndSay(n - 1)
	for i := 0; i < len(inp); i++ {
		count++
		if i == len(inp)-1 || inp[i] != inp[i+1] {
			res.WriteString(strconv.Itoa(count))
			res.WriteByte(inp[i])
			// res += strconv.Itoa(count) + string(inp[i])
			count = 0
		}
	}
	memo[n] = res.String()
	return res.String()
}

func main() {
	println("Started...")
	test := 1
	fmt.Printf("Testing: %v\n", test)
	result := countAndSay(test)
	fmt.Printf("Yields : %v\n", result)

	test = 2
	fmt.Printf("Testing: %v\n", test)
	result = countAndSay(test)
	fmt.Printf("Yields : %v\n", result)

	test = 10
	fmt.Printf("Testing: %v\n", test)
	result = countAndSay(test)
	fmt.Printf("Yields : %v\n", result)

	test = 30
	fmt.Printf("Testing: %v\n", test)
	result = countAndSay(test)
	fmt.Printf("Yields : %v\n", result)
}
