package answer

import (
	"fmt"

	. "leetcode/utils"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ld, rd := minDepth(root.Left), minDepth(root.Right)
	if ld == 0 || rd == 0 {
		return ld + rd + 1
	}
	if ld > rd {
		return rd + 1
	} else {
		return ld + 1
	}
}

func (sol *Solution) Title111() {
	root := CreateTree([]int{3, 9, 20, -1, -1, 15, 7})
	PrintTree(root)
	result := minDepth(root)
	fmt.Println(result)
}
