package answer

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)

	tlist := make([]int, 0)
	push := func(i int) {
		// 保证队列中下标对应的元素是递减的
		for len(tlist) > 0 && nums[i] >= nums[tlist[len(tlist)-1]] {
			tlist = tlist[:len(tlist)-1]
		}
		tlist = append(tlist, i)
	}
	for i := 0; i < k; i++ {
		// 初始化窗口队列
		push(i)
	}

	res[0] = nums[tlist[0]]
	for i := k; i < len(nums); i++ {
		push(i)
		// 如果队头的元素已经不在窗口内，则移除
		for tlist[0] <= i-k {
			tlist = tlist[1:]
		}
		res[i-k+1] = nums[tlist[0]]
	}

	return res
}

func (sol *Solution) Title239() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}
