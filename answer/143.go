package answer

import (
	. "leetcode/structures"
)

func reorderList(head *ListNode) {
	nodeNum := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		nodeNum++
	}
	tmp := head
	for i := 0; i < nodeNum/2; i++ {
		tmp = tmp.Next
	}
	var tail *ListNode
	curNode := tmp
	for tmp.Next != nil {
		nextNode := tmp.Next
		if nextNode.Next == nil {
			tail = nextNode
		}
		tmp.Next = nextNode.Next
		nextNode.Next = curNode
		curNode = nextNode
	}
	for head != tmp && tail != tmp {
		nextNode := head.Next
		head.Next = tail
		tail = tail.Next
		head.Next.Next = nextNode
		head = nextNode
	}
}

func (sol *Solution) Title143() {
	head := CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	reorderList(head)
	PrintList(head)
}
