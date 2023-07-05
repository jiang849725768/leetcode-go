package answer

import (
	"fmt"

	. "leetcode/utils"
)

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	const (
		left = iota
		mid
		right
	)

	var dfs func(root *TreeNode, flag int)
	dfs = func(root *TreeNode, flag int) {
		isLeaf := true
		if root.Left != nil {
			isLeaf = false
			dfs(root.Left, left)
		}
		if root.Right != nil {
			isLeaf = false
			dfs(root.Right, right)
		}
		if isLeaf && flag == left {
			sum += root.Val
		}
	}
	dfs(root, mid)

	return sum
}

func (sol *Solution) Title404() {
	root := CreateTree([]int{3, 9, 20, -1, -1, 15, 7})
	result := sumOfLeftLeaves(root)
	fmt.Println(result)
}
