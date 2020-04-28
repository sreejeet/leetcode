package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	slice := strings.Split(strings.TrimRight(s, " "), " ")
	// fmt.Printf("%v | %v\n", slice, slice[len(slice)-1])
	return len(slice[len(slice)-1])
	return 0
}

func main() {
	test := ""
	fmt.Printf("Testing: %v\n", test)
	result := lengthOfLastWord(test)
	fmt.Printf("Yields : %v\n", result)

	test = "asd asdf"
	fmt.Printf("Testing: %v\n", test)
	result = lengthOfLastWord(test)
	fmt.Printf("Yields : %v\n", result)

	test = " asd asdf asd asd a "
	fmt.Printf("Testing: %v\n", test)
	result = lengthOfLastWord(test)
	fmt.Printf("Yields : %v\n", result)
}
