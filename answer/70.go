package answer

import (
	"fmt"
)

func climbStairs(n int) int {
	dp := make([]int, n+1)
	steps := []int{1, 2}
	dp[0] = 1

	for i := 0; i < n+1; i++ {
		for j := range steps {
			if steps[j] <= i {
				dp[i] += dp[i-steps[j]]
			}
		}
	}

	return dp[n]
}

func (sol *Solution) Title70() {
	n := 2
	result := climbStairs(n)
	fmt.Println(result)
}
