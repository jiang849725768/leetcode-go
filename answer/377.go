package answer

import (
	"fmt"
)

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i <= target; i++ {
		for j := range nums {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}

func (sol *Solution) Title377() {
	nums := []int{1, 2, 3}
	target := 4
	result := combinationSum4(nums, target)
	fmt.Println(result)
}
