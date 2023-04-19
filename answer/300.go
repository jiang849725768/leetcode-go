package answer

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	// 78912345 789 189 129 123 1234 12345
	for _, num := range nums {
		i, j := 0, res
		// 二分查找
		for i < j {
			mid := (i + j) / 2
			if nums[mid] < num {
				i = mid + 1
			} else {
				j = mid
			}
		}
		// 原地修改nums以维护一个递增最小序列数组
		nums[i] = num
		if res == j {
			res++
		}
	}
	return res
}

func (sol *Solution) Title300() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
