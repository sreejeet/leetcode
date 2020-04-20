package main

import (
	"fmt"
)

func romanToInt(s string) int {
	roman := map[rune]uint16{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res uint16 = 0
	last := uint16(0)
	for i := len(s) - 1; i >= 0; i-- {
		current := roman[rune(s[i])]
		if last > current {
			res -= current
		} else {
			res += current
		}
		last = current
	}
	// for _, r := range s {
	// print(i, " ", roman[r], "\n")
	// }
	return int(res)
}

func main() {
	test := "MCMXCIV"
	result := romanToInt(test)
	fmt.Println(test)
	fmt.Println(result, "\n")

	test = "LVIII"
	result = romanToInt(test)
	fmt.Println(test)
	fmt.Println(result, "\n")

	test = "MMMCMXCIX"
	result = romanToInt(test)
	fmt.Println(test)
	fmt.Println(result, "\n")
}
