package answer

import (
	"fmt"
)

func trap(height []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	store := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if height[left] < height[right] {
			store += leftMax - height[left]
			left++
		} else {
			store += rightMax - height[right]
			right--
		}
	}

	return store
}

func (sol *Solution) Title42() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
