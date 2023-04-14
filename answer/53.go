package answer

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	var value, res int
	res = nums[0]
	for i := range nums {
		if value < 0 && nums[i] > value {
			value = nums[i]
		} else {
			value += nums[i]
		}
		if value > res {
			res = value
		}
	}
	return res
}

func (sol *Solution) Title53() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
