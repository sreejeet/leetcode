package main

import "fmt"

func lengthOfLongestSubstring(s string) (max int) {
	if len(s) < 2 {
		return len(s)
	}
	max = 1
	hash := map[byte]int{s[0]: 0}
	x := 0
	for l := 1; l < len(s); l++ {
		// println(s[x:l], x, l, max)
		if ix, exist := hash[s[l]]; exist && ix >= x {
			// println(string(s[l]), "at", ix)
			x = ix + 1
		}
		if l-x >= max {
			max = l - x + 1
		}
		hash[s[l]] = l
	}
	return
}

func main() {
	test := "aaaaaaaasd"
	println("Testing", test)
	result := lengthOfLongestSubstring(test)
	fmt.Printf("Yields %v\n", result)
}

/*
""
"a"
"aa"
"ab"
"aaa"
"abb"
"bba"
"abc"
"asdasdd"
"aaaaaaaasd"
"abcabcbb"
"asdfggfdsa"
"asdffghjkl"
*/
