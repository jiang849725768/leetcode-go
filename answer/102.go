package answer

import (
	. "leetcode/structures"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	layer := []*TreeNode{root}
	for len(layer) > 0 {
		var level []int
		nextLayer := []*TreeNode{}
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
