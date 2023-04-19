package answer

import (
	. "leetcode/utils"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mList := func(left, right *ListNode) *ListNode {
		fakeHead := &ListNode{}
		cur := fakeHead
		for left != nil && right != nil {
			if left.Val < right.Val {
				cur.Next = left
				left = left.Next
			} else {
				cur.Next = right
				right = right.Next
			}
			cur = cur.Next
		}
		if left != nil {
			cur.Next = left
		}
		if right != nil {
			cur.Next = right
		}
		return fakeHead.Next
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	return mList(left, right)
}

func (sol *Solution) Title148() {
	head := CreateList([]int{4, 2, 1, 3})
	PrintList(head)
	PrintList(sortList(head))
}
