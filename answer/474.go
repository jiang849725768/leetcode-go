package answer

import (
	"fmt"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range strs {
		zeroNum := strings.Count(strs[i], "0")
		oneNum := strings.Count(strs[i], "1")

		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				cnt := dp[j-zeroNum][k-oneNum] + 1
				if dp[j][k] < cnt {
					dp[j][k] = cnt
				}
			}
		}
	}

	return dp[m][n]
}

func (sol *Solution) Title474() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	result := findMaxForm(strs, m, n)
	fmt.Println(result)
}
