package answer

import (
	. "leetcode/utils"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			root := &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1], inorder[:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
			return root
		}
	}
	return nil
}

func (sol *Solution) Title105() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	PrintTree(root)
}
