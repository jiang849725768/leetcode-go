package answer

import (
	"fmt"

	. "leetcode/utils"
)

func isValidBST(root *TreeNode) bool {
	var check func(root *TreeNode, min, max int) bool
	check = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		return check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
	}

	return check(root, -1<<31-1, 1<<31)
}

func (sol *Solution) Title98() {
	root := CreateTree([]int{2, 1, 3})
	PrintTree(root)
	fmt.Println(isValidBST(root))
	root = CreateTree([]int{5, 1, 4, -1, -1, 3, 6})
	PrintTree(root)
	fmt.Println(isValidBST(root))
}
