package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	m := x / 2
	for s, e := 1, x; s < e-1; m = (s + e) / 2 {
		if m*m > x {
			e = m
			continue
		}
		s = m
	}
	return m
}

func main() {
	test := 2837658
	println("IN :", test)
	result := mySqrt(test)
	println("OUT:", result)

	test = 0
	println("IN :", test)
	result = mySqrt(test)
	println("OUT:", result)

	test = 1
	println("IN :", test)
	result = mySqrt(test)
	println("OUT:", result)
}

/*
0
1
2
3
123
2345
34567
456789
5678901
67890123
789012345
*/
