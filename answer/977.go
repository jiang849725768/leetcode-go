package answer

import "fmt"

func sortedSquares(nums []int) []int {
	head, tail := 0, len(nums)-1
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if a, b := nums[head]*nums[head], nums[tail]*nums[tail]; a > b {
			res[i] = a
			head++
		} else {
			res[i] = b
			tail--
		}
	}
	return nums
}

func (sol *Solution) Title977() {
	nums := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(nums)
	fmt.Println(result)
	nums = []int{-7, -3, 2, 3, 11}
	result = sortedSquares(nums)
	fmt.Println(result)
}
