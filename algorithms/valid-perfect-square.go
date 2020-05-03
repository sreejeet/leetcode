package main

func isPerfectSquare(x int) bool {
	if x < 2 {
		return true
	}
	s, e := 1, x
	m := x / 2
	for s < e-1 {
		m = (s + e) / 2
		if m*m == x {
			return true
		}
		if m*m > x {
			e = m
		} else {
			s = m
		}
	}
	return false
}

func main() {
	test := 81
	println("IN :", test)
	result := isPerfectSquare(test)
	println("OUT:", result)
}
