package answer

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := range coins {
		for j := coins[i]; j < amount+1; j++ {
			if dp[j] > dp[j-coins[i]]+1 {
				dp[j] = dp[j-coins[i]] + 1
			}
		}
	}

	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1
}

func (sol *Solution) Title322() {
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	fmt.Println(result)
}
