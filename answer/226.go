package answer

import (
	. "leetcode/utils"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}

func (sol *Solution) Title226() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	PrintTree(root)
	result := invertTree(root)
	PrintTree(result)
}
