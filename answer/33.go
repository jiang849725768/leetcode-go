package answer

import (
	"fmt"
)

func search(nums []int, target int) int {
	bg, ed := 0, len(nums)-1
	for bg <= ed {
		if nums[bg] < nums[ed] && target >= nums[bg] && target <= nums[ed] {
		}
	}
	return -1
}

func (sol *Solution) Title33() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
