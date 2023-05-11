package answer

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	var qsort func(left, right int)
	qsort = func(left, right int) {
		start, end := left, right
		if start >= end {
			return
		}
		for start < end {
			for start < end && nums[start] < nums[end] {
				end--
			}
			if start < end {
				nums[start], nums[end] = nums[end], nums[start]
				start++
			}
			for start < end && nums[start] < nums[end] {
				start++
			}
			if start < end {
				nums[start], nums[end] = nums[end], nums[start]
				end--
			}
		}
		if start > len(nums)-k {
			qsort(left, start-1)
		} else if start < len(nums)-k {
			qsort(end+1, right)
		}
	}
	qsort(0, len(nums)-1)

	return nums[len(nums)-k]
}

func (sol *Solution) Title215() {
	values := []int{5, 2, 4, 1, 3, 6, 0}
	k := 4
	fmt.Println(findKthLargest(values, k))
}
