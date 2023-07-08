package answer

import (
	"fmt"
)

func findTargetSumWays(nums []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if (target+sum)%2 == 1 || abs(target) > sum {
		return 0
	}

	dp := make([]int, (target+sum)/2+1)
	// dp[0] means take no element, so it is 1
	dp[0] = 1

	for i := range nums {
		for j := (target + sum) / 2; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[(target+sum)/2]
}

func (sol *Solution) Title494() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	result := findTargetSumWays(nums, target)
	fmt.Println(result)
}
