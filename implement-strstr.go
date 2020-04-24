package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && len(haystack)-i >= len(needle) {
			foundNeedle := true
			for j := 0; j < len(needle) && j < len(haystack); j++ {
				if needle[j] != haystack[i+j] {
					foundNeedle = false
					break
				}
			}
			if foundNeedle {
				return i
			}
		}
	}
	return -1
}

func main() {
	test := "233415"
	fmt.Printf("Testing: %v\n", test)
	result := strStr(test, "1")
	fmt.Printf("Yields : %v\n", result)

}

/*
class Solution {
public:
    int strStr(string haystack, string needle) {

    }
};*/
