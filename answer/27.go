package answer

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i := 0
	for j := range nums {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func (sol *Solution) Title27() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(nums[:removeElement(nums, val)])
}
