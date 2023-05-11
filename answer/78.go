package answer

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for len(nums) > 0 {
		lenRes := len(res)
		for i := 0; i < lenRes; i++ {
			// 直接append会导致res[i]和res[lenRes+i]指向同一个底层数组
			tmp := make([]int, len(res[i]))
			copy(tmp, res[i])
			tmp = append(tmp, nums[0])
			res = append(res, tmp)
		}
		nums = nums[1:]
	}

	return res
}

func (sol *Solution) Title78() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
