package answer

import (
	"fmt"
)

func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := range stones {
		sum += stones[i]
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := range stones {
		for j := target; j >= stones[i]; j-- {
			if dp[j] < dp[j-stones[i]]+stones[i] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	if 2*dp[target] > sum {
		return 2*dp[target] - sum
	} else {
		return sum - 2*dp[target]
	}
}

func (sol *Solution) Title1049() {
	stones := []int{2, 7, 4, 1, 8, 1}
	result := lastStoneWeightII(stones)
	fmt.Println(result)
}
