package main

import (
	"fmt"
)

func isValid(s string) bool {
	parentMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, char := range s {
		if _, ok := parentMap[char]; ok {
			stack = append(stack, char)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		v, _ := parentMap[stack[len(stack)-1]]
		if char == v {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return len(stack) == 0
}

func main() {
	test := "{"
	fmt.Printf("Testing  :%v\n", test)
	result := isValid(test)
	fmt.Printf("Yields   :%v\n", result)

	test = "{[(]})"
	fmt.Printf("Testing  :%v\n", test)
	result = isValid(test)
	fmt.Printf("Yields   :%v\n", result)

	test = "{}{}(){()}"
	fmt.Printf("Testing  :%v\n", test)
	result = isValid(test)
	fmt.Printf("Yields   :%v\n", result)

	test = "[(]{)}"
	fmt.Printf("Testing  :%v\n", test)
	result = isValid(test)
	fmt.Printf("Yields   :%v\n", result)
}
