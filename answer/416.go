package answer

import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	dp := make([]int, sum/2+1)
	for i := range nums {
		for j := sum / 2; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}
		}
	}

	return dp[sum/2] == sum/2
}

func (sol *Solution) Title416() {
	nums := []int{1, 5, 11, 5}
	result := canPartition(nums)
	fmt.Println(result)
}
