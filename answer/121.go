package answer

import (
	"fmt"
)

func maxProfit(prices []int) int {
	minValue := prices[0]
	res := 0
	for i := range prices {
		curProfit := prices[i] - minValue
		if prices[i] < minValue {
			minValue = prices[i]
		}
		if curProfit > res {
			res = curProfit
		}
	}

	return res
}

func (sol *Solution) Title121() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
