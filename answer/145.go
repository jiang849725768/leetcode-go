package answer

import (
	"fmt"

	. "leetcode/utils"
)

func postorderTraversal(root *TreeNode) []int {
	//迭代
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	nlist := []*TreeNode{root}
	for len(nlist) != 0 {
		cur := nlist[len(nlist)-1]
		nlist = nlist[:len(nlist)-1]
		res = append(res, cur.Val)
		if cur.Left != nil {
			nlist = append(nlist, cur.Left)
		}
		if cur.Right != nil {
			nlist = append(nlist, cur.Right)
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	//递归
	//res := make([]int, 0)
	//
	//var dfs func(root *TreeNode)
	//dfs = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//	dfs(root.Left)
	//	dfs(root.Right)
	//	res = append(res, root.Val)
	//}
	//dfs(root)

	return res
}

func (sol *Solution) Title145() {
	root := CreateTree([]int{1, 2, 3, 4, 5, 6, 7})
	PrintTree(root)
	result := postorderTraversal(root)
	fmt.Println(result)
}
