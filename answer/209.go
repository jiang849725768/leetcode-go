package answer

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	sum := 0
	res := len(nums) + 1
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			winLen := end - start + 1
			if winLen < res {
				res = winLen
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func (sol *Solution) Title209() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	result := minSubArrayLen(target, nums)
	fmt.Println(result)
	target = 4
	nums = []int{1, 4, 4}
	result = minSubArrayLen(target, nums)
	fmt.Println(result)
	target = 11
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	result = minSubArrayLen(target, nums)
	fmt.Println(result)
}
