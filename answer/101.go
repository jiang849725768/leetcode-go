package answer

import (
	"fmt"

	. "leetcode/utils"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var same func(a, b *TreeNode) bool
	same = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		} else if a == nil || b == nil {
			return false
		}

		return a.Val == b.Val && same(a.Left, b.Right) && same(a.Right, b.Left)
	}

	return same(root.Left, root.Right)
}

func (sol *Solution) Title101() {
	root := CreateTree([]int{1, 2, 2, 3, 4, 4, 3})
	PrintTree(root)
	result := isSymmetric(root)
	fmt.Println(result)
}
