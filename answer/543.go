package answer

import (
	"fmt"

	. "leetcode/utils"
)

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		r, l := depth(node.Left), depth(node.Right)
		if r+l > res {
			res = r + l
		}
		if r > l {
			return r + 1
		}
		return l + 1
	}
	depth(root)

	return res
}

func (sol *Solution) Title543() {
	root := CreateTree([]int{1, 2, 3, 4, 5})
	result := diameterOfBinaryTree(root)
	fmt.Println(result)
}
