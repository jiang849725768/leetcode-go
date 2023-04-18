package answer

import (
	"fmt"
)

// 1432 2134 2314
// 123 132 213 231 312 321
func nextPermutation(nums []int) {
	for j := len(nums) - 1; j > 0; j-- {
		if nums[j] > nums[j-1] {
			for k := len(nums) - 1; k >= j; k-- {
				if nums[k] > nums[j-1] {
					nums[k], nums[j-1] = nums[j-1], nums[k]
					break
				}
			}
			for k := j; k < (j+len(nums))/2; k++ {
				nums[k], nums[j+len(nums)-k-1] = nums[j+len(nums)-k-1], nums[k]
			}
			return
		}
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func (sol *Solution) Title31() {
	nums := []int{1, 4, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}
