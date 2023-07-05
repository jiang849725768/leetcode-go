package answer

import (
	"fmt"

	. "leetcode/utils"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	countDepth := func(root *TreeNode) int {
		depth := 0
		for root != nil {
			depth++
			root = root.Left
		}
		return depth
	}
	for {
		res++
		ldep, rdep := countDepth(root.Left), countDepth(root.Right)
		if ldep == 0 || rdep == 0 {
			res = res + ldep + rdep
			break
		}
		if ldep == rdep {
			res += 1<<ldep - 1
			root = root.Right
		} else if ldep > rdep {
			res += 1<<rdep - 1
			root = root.Left
		}
	}

	return res
}

func (sol *Solution) Title222() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6})
	result := countNodes(root)
	fmt.Println(result)
}
