package answer

import (
	"fmt"
	"strconv"

	. "leetcode/utils"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	var readTree func(root *TreeNode, part string)
	readTree = func(root *TreeNode, part string) {
		leaf := true
		part += strconv.Itoa(root.Val)
		if root.Left != nil {
			leaf = false
			readTree(root.Left, part+"->")
		}
		if root.Right != nil {
			leaf = false
			readTree(root.Right, part+"->")
		}
		if leaf {
			res = append(res, part)
		}
	}
	readTree(root, "")
	return res
}

func (sol *Solution) Title257() {
	root := CreateTree([]int{1, 2, 3, -1, 5})
	result := binaryTreePaths(root)
	fmt.Println(result)
}
