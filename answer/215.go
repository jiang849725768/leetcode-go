package answer

import (
	"fmt"
)

// QuickSort
func findKthLargest(nums []int, k int) int {
	var TopK func(left, right int)
	TopK = func(left, right int) {
		if left >= right {
			return
		}
		bg, ed := left, right
		for left < right {
			for nums[right] > nums[bg] && left < right {
				right--
			}
			for nums[left] <= nums[bg] && left < right {
				left++
			}
			nums[left], nums[right] = nums[right], nums[left]
		}
		nums[bg], nums[left] = nums[left], nums[bg]
		switch {
		case right == len(nums)-k:
			return
		case right < len(nums)-k:
			TopK(right+1, ed)
		case right > len(nums)-k:
			TopK(bg, left-1)
		default:
			panic("Impossible")
		}
	}
	TopK(0, len(nums)-1)
	return nums[len(nums)-k]
}

func (sol *Solution) Title215() {
	values := []int{5, 2, 4, 1, 3, 6, 0}
	k := 4
	fmt.Println(findKthLargest(values, k))
}
