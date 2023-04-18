package utils

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CreateTree create a tree with int ptr slice
// value -1 is used to represent null
func CreateTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	head := &TreeNode{Val: nums[0]}
	nodes := make([]*TreeNode, len(nums))
	nodes[0] = head
	for i := 1; i < len(nums); i++ {
		// value -1 is used to represent null
		if nums[i] == -1 {
			nodes[i] = nil
			continue
		}
		nodes[i] = &TreeNode{Val: nums[i]}
	}
	for i := 0; i < len(nums)/2; i++ {
		if nodes[i] == nil {
			continue
		}
		nodes[i].Left = nodes[2*i+1]
		if 2*i+2 < len(nums) {
			nodes[i].Right = nodes[2*i+2]
		}
	}
	return head
}

// PrintTree print value of tree in tree mode
func PrintTree(head *TreeNode) {
	nodeTree := make([][]*TreeNode, 0)
	layer := []*TreeNode{head}
	validNode := 1
	for validNode != 0 {
		nodeTree = append(nodeTree, layer)
		validNode = 0
		nextLayer := make([]*TreeNode, 0)
		for _, node := range layer {
			if node == nil {
				nextLayer = append(nextLayer, nil, nil)
				continue
			}
			nextLayer = append(nextLayer, node.Left, node.Right)
			if node.Left != nil {
				validNode++
			}
			if node.Right != nil {
				validNode++
			}
		}
		layer = nextLayer
	}

	layerLength := 5*len(nodeTree[len(nodeTree)-1]) - 1
	for i := 0; i < len(nodeTree); i++ {
		spaceNum := (layerLength - 5*len(nodeTree[i]) + 1) / 2
		fmt.Print(strings.Repeat(" ", spaceNum))
		layer := nodeTree[i]
		for _, node := range layer {
			if node == nil {
				fmt.Printf("null ")
			} else {
				fmt.Printf("%-4d ", node.Val)
			}
		}
		fmt.Printf("\n")
	}
}
