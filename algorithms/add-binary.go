package main

import "strconv"

func addBinary(a string, b string) string {
	res := ""
	s, lenA, lenB := 0, len(a), len(b)
	for ptr := 1; ptr <= lenA || ptr <= lenB; ptr++ {
		if ptr <= lenA && a[lenA-ptr] == '1' {
			s++
		}
		if ptr <= lenB && b[lenB-ptr] == '1' {
			s++
		}
		res = strconv.Itoa(s%2) + res
		s /= 2
	}
	if s == 1 {
		res = "1" + res
	}
	return res
}

func main() {
	// test1 := "1"
	// test2 := "1"
	test1 := "11111101010"
	test2 := "110101011"
	println(test1, test2)
	println()
	result := addBinary(test1, test2)
	println()
	println(result)
}

/*
"11111101010"
"110101011"
111010010110"
*/
