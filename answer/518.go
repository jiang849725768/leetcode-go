package answer

import (
	"fmt"
)

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := range coins {
		for j := coins[i]; j < amount+1; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}

func (sol *Solution) Title518() {
	amount := 5
	coins := []int{1, 2, 5}
	result := change(amount, coins)
	fmt.Println(result)
}
