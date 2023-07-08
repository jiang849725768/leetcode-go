package answer

import (
	"fmt"
)

func numSquares(n int) int {
	x := n
	for {
		nextX := (x + n/x) / 2
		if nextX >= x {
			break
		}
		x = nextX
	}
	fmt.Println("x is ", x)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1 << 15
	}
	for i := 1; i <= x; i++ {
		val := i * i
		for j := val; j < n+1; j++ {
			if dp[j] > dp[j-val]+1 {
				dp[j] = dp[j-val] + 1
			}

		}
		fmt.Println(dp)
	}

	return dp[n]
}

func (sol *Solution) Title279() {
	n := 3
	result := numSquares(n)
	fmt.Println(result)
}
