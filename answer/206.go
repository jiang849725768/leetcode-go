package answer

import (
	. "leetcode/utils"
)

func reverseList(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	cur := head
	var pre *ListNode
	for fakeHead.Next != nil {
		fakeHead.Next = cur.Next
		cur.Next = pre
		pre = cur
		if fakeHead.Next != nil {
			cur = fakeHead.Next
		}
	}
	return cur
}

func (sol *Solution) Title206() {
	head := CreateList([]int{1, 2, 3, 4, 5})
	result := reverseList(head)
	PrintList(result)
}
