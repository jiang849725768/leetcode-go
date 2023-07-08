package answer

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s)+1; i++ {
		for j := range wordDict {
			lword := len(wordDict[j])
			if i >= lword && s[i-lword:i] == wordDict[j] {
				dp[i] = dp[i] || dp[i-lword]
			}
		}
		fmt.Println(dp)
	}

	return dp[len(s)]
}

func (sol *Solution) Title139() {
	s := "dogs"
	wordDict := []string{"dog", "s", "gs"}
	result := wordBreak(s, wordDict)
	fmt.Println(result)
}
