package answer

import (
	. "leetcode/utils"
)

func removeElements(head *ListNode, val int) *ListNode {
	fakeHead := &ListNode{Next: head}
	for tmp := fakeHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return fakeHead.Next
}

func (sol *Solution) Title203() {
	head := CreateList([]int{1, 2, 6, 3, 4, 5, 6})
	val := 6
	result := removeElements(head, val)
	PrintList(result)
}
