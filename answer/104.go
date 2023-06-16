package answer

import (
	"fmt"

	. "leetcode/utils"
)

func maxDepth(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			if res < depth {
				res = depth
			}
			return
		}
		depth++
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}
	dfs(root, 0)

	return res
}

func (sol *Solution) Title104() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	PrintTree(root)
	result := maxDepth(root)
	fmt.Println(result)
}
