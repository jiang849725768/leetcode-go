package answer

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var res [][]int
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums)-1 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	dfs(0)

	return res
}

func (sol *Solution) Title46() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
