package answer

import (
	"fmt"

	. "leetcode/utils"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	layer := []*TreeNode{root}
	for len(layer) > 0 {
		var level []int
		var nextLayer []*TreeNode
		for i := range layer {
			if layer[i].Left != nil {
				nextLayer = append(nextLayer, layer[i].Left)
			}
			if layer[i].Right != nil {
				nextLayer = append(nextLayer, layer[i].Right)
			}
			level = append(level, layer[i].Val)
		}
		layer = nextLayer
		res = append(res, level)
	}
	return res
}

func (sol *Solution) Title102() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	PrintTree(root)
	result := levelOrder(root)
	fmt.Println(result)
}
