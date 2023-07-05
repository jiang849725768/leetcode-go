package answer

import (
	"fmt"

	. "leetcode/utils"
)

func findBottomLeftValue(root *TreeNode) int {
	tlist := []*TreeNode{root}
	for len(tlist) != 0 {
		nlist := make([]*TreeNode, 0)
		for i := range tlist {
			node := tlist[i]
			if node.Left != nil {
				nlist = append(nlist, node.Left)
			}
			if node.Right != nil {
				nlist = append(nlist, node.Right)
			}
		}
		if len(nlist) == 0 {
			return tlist[0].Val
		}
		tlist = nlist
	}

	return 0
}

func (sol *Solution) Title513() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	result := findBottomLeftValue(root)
	fmt.Println(result)
}
