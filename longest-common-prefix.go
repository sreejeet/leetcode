package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var lastPrefix string = ""
	for v := 0; ; v++ {
		for h := 0; h <= len(strs)-2; h++ {
			if len(strs[h]) < v || len(strs[h+1]) < v || strs[h][:v] != strs[h+1][:v] {
				return lastPrefix
			}
		}
		lastPrefix = strs[0][:v]
	}
	return lastPrefix
}

func main() {
	test := []string{"asd", "bsd"}
	fmt.Println("Testing:", test)
	result := longestCommonPrefix(test)
	fmt.Println("Yields :", result, "\n")

	test = []string{"a", "b"}
	fmt.Println("Testing:", test)
	result = longestCommonPrefix(test)
	fmt.Println("Yields :", result, "\n")

	test = []string{"xc", "xc", "", "xc"}
	fmt.Println("Testing:", test)
	result = longestCommonPrefix(test)
	fmt.Println("Yields :", result, "\n")

	test = []string{"abcd", "abc"}
	fmt.Println("Testing:", test)
	result = longestCommonPrefix(test)
	fmt.Println("Yields :", result, "\n")

	test = []string{"abc", "abcd"}
	fmt.Println("Testing:", test)
	result = longestCommonPrefix(test)
	fmt.Println("Yields :", result, "\n")

	test = []string{"abcd", "abcde", "abcd"}
	fmt.Println("Testing:", test)
	result = longestCommonPrefix(test)
	fmt.Println("Yields :", result, "\n")
}
