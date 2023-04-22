package answer

import (
	. "leetcode/utils"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{
		Next: head,
	}
	slow, fast := fakeHead, fakeHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return fakeHead.Next
}

func (sol *Solution) Title19() {
	head := CreateList([]int{1, 2, 3, 4, 5})
	PrintList(head)
	PrintList(removeNthFromEnd(head, 2))
}
