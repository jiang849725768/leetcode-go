package answer

import (
	"fmt"

	. "leetcode/utils"
)

func maxPathSum(root *TreeNode) int {
	res := root.Val
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		left, right := 0, 0
		if root.Left != nil {
			left = max(helper(root.Left), 0)
		}
		if root.Right != nil {
			right = max(helper(root.Right), 0)
		}
		if left+right+root.Val > res {
			res = left + right + root.Val
		}
		return max(left, right) + root.Val
	}
	helper(root)
	return res
}

func (sol *Solution) Title124() {
	root := CreateTree([]int{-10, 9, 20, -1, -1, 15, 7})
	PrintTree(root)
	fmt.Println(maxPathSum(root))
}
