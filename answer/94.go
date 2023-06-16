package answer

import (
	"fmt"

	. "leetcode/utils"
)

func inorderTraversal(root *TreeNode) []int {
	// 迭代
	//var res []int
	//var stack []*TreeNode
	//for root != nil || len(stack) > 0 {
	//	if root != nil {
	//		stack = append(stack, root) // 入栈
	//		root = root.Left
	//	} else {
	//		root = stack[len(stack)-1] // 出栈
	//		stack = stack[:len(stack)-1]
	//		res = append(res, root.Val)
	//		root = root.Right
	//	}
	//}

	// 递归
	res := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)

	return res
}

func (sol *Solution) Title94() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(inorderTraversal(root))
}
