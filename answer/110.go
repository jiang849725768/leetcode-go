package answer

import (
	"fmt"

	. "leetcode/utils"
)

func isBalanced(root *TreeNode) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh := height(root.Left)
		if lh == -1 {
			return -1
		}
		rh := height(root.Right)
		if rh == -1 || abs(lh-rh) > 1 {
			return -1
		}
		return max(lh, rh) + 1
	}

	return height(root) >= 0
}

func (sol *Solution) Title110() {
	root := CreateTree([]int{3, 9, 20, -1, -1, 15, 7})
	result := isBalanced(root)
	fmt.Println(result)
}
