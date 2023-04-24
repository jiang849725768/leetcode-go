package answer

import (
	"fmt"

	. "leetcode/utils"
)

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root) // 入栈
			root = root.Left
		} else {
			root = stack[len(stack)-1] // 出栈
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}

func (sol *Solution) Title94() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(inorderTraversal(root))
}
