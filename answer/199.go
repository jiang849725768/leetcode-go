package answer

import (
	"fmt"

	. "leetcode/utils"
)

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		res = append(res, queue[len(queue)-1].Val)
		tmp := make([]*TreeNode, 0)
		for i := range queue {
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		queue = tmp
	}

	return res
}

func (sol *Solution) Title199() {
	root := CreateTree([]int{1, 2, 3, 4})
	PrintTree(root)
	fmt.Println(rightSideView(root))
}
