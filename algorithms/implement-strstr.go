package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := range haystack {
		if haystack[i] == needle[0] &&
			len(haystack)-i >= len(needle) &&
			haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	test := "23345"
	fmt.Printf("Testing: %v\n", test)
	result := strStr(test, "1")
	fmt.Printf("Yields : %v\n", result)

	test = "1"
	fmt.Printf("Testing: %v\n", test)
	result = strStr(test, "1")
	fmt.Printf("Yields : %v\n", result)

	test = "231345"
	fmt.Printf("Testing: %v\n", test)
	result = strStr(test, "1")
	fmt.Printf("Yields : %v\n", result)

	test = "23345a"
	fmt.Printf("Testing: %v\n", test)
	result = strStr(test, "a")
	fmt.Printf("Yields : %v\n", result)

}
