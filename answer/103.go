package answer

import (
	"fmt"
	. "leetcode/structures"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	layer := 0
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		nextQueue := make([]*TreeNode, 0)
		tmp := make([]int, 0)
		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		if layer%2 == 0 {
			for _, node := range queue {
				tmp = append(tmp, node.Val)
			}
		} else {
			for i := len(queue) - 1; i >= 0; i-- {
				tmp = append(tmp, queue[i].Val)
			}
		}
		layer++
		queue = nextQueue
		res = append(res, tmp)
	}

	return res
}

func (sol *Solution) Title103() {
	root := CreateTree([]int{3, 9, 20, -1, -1, 15, 7})
	PrintTree(root)
	res := zigzagLevelOrder(root)
	for _, ints := range res {
		for _, i := range ints {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\n")
	}

}
