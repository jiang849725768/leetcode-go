package answer

import (
	. "leetcode/structures"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fakeHead := &ListNode{Val: head.Val - 1, Next: head}
	dupValue := head.Val - 1
	pre := fakeHead
	cur := head
	for {
		for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
			dupValue = cur.Val
			for cur != nil && cur.Val == dupValue {
				cur = cur.Next
			}
			pre.Next = cur
		}
		if cur != nil {
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}

	return fakeHead.Next
}

func (sol *Solution) Title82() {
	head := CreateList([]int{1, 2, 3, 3, 4, 4, 5})
	PrintList(deleteDuplicates(head))
	head = CreateList([]int{1, 1, 1, 2, 3})
	PrintList(deleteDuplicates(head))
}
