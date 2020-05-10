package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	buy := prices[0]
	for _, sell := range prices {
		if sell-buy > profit {
			profit = sell - buy
		}
		if sell < buy {
			buy = sell
		}
	}
	return profit
}

func main() {
	test := []int{1, 6, 5, 4, 3, 2, 1, 0, -1}
	fmt.Println("Testing", test)
	result := maxProfit(test)
	fmt.Printf("Yields %v\n", result)
}

/*
[7,1,5,3,6,4]
[7,6,4,3,1]
[1, 5, 7, 2, 3, 24, 7, 9, 2, 4, 62, 46, 35, 84, 69, 46, 246, 2, 4, 62, 7, 458, 46, 84, 82, 4, 6357, 457, 468, 3, 62, 42, 42, 34, 236, 346, 457, 468, 58, 785]
*/
