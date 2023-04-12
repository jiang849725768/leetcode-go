package answer

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var quickSort func(left, right int)
	quickSort = func(left, right int) {
		if left >= right {
			return
		}
		pt := rand.Intn(right-left+1) + left
		nums[left], nums[pt] = nums[pt], nums[left]
		bg, ed := left, right
		for left < right {
			for nums[left] < nums[right] && left < right {
				right--
			}
			nums[left], nums[right] = nums[right], nums[left]
			left++
			for nums[left] < nums[right] && left < right {
				left++
			}
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
		quickSort(bg, left-1)
		quickSort(right+1, ed)
	}
	quickSort(0, len(nums)-1)

	return nums
}

func (sol *Solution) Title912() {
	// nums := []int{5, 1, 1, 2, 0, 0}
	nums := []int{3, -1}
	nums = sortArray(nums)
	fmt.Println(nums)
}
