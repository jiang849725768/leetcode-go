package answer

import (
	"fmt"

	. "leetcode/utils"
)

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	nlist := []*TreeNode{root}
	if root == nil {
		return []int{}
	}

	for len(nlist) != 0 {
		cur := nlist[len(nlist)-1]
		nlist = nlist[:len(nlist)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			nlist = append(nlist, cur.Right)
		}
		if cur.Left != nil {
			nlist = append(nlist, cur.Left)
		}
	}

	return res
}

func (sol *Solution) Title144() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	PrintTree(root)
	result := preorderTraversal(root)
	fmt.Println(result)
}
