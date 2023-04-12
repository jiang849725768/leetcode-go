package answer

import (
	. "leetcode/structures"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	fakeHead := &ListNode{Next: head}
	preNode := fakeHead
	num := 1
	for num < left {
		num++
		preNode = preNode.Next
	}
	startNode := preNode.Next
	for num < right {
		num++
		nextNode := startNode.Next
		startNode.Next = nextNode.Next
		nextNode.Next = preNode.Next
		preNode.Next = nextNode
	}

	return fakeHead.Next
}

func (sol *Solution) Title92() {
	head := CreateList([]int{1, 2, 3, 4, 5})
	left, right := 2, 4
	head = reverseBetween(head, left, right)
	PrintList(head)
}
