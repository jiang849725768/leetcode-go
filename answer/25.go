package answer

import (
	. "leetcode/utils"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	fakeHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	lastTurntail := fakeHead
	for node != nil {
		count := 0
		turnHead := node
		for count != k {
			if node == nil {
				break
			}
			node = node.Next
			count++
		}
		if count != k {
			break
		}
		nextNode, lastNode := node, turnHead
		for lastNode.Next != node {
			tmp := lastNode.Next
			lastNode.Next = nextNode
			nextNode, lastNode = lastNode, tmp
		}
		lastNode.Next = nextNode
		lastTurntail.Next = lastNode
		lastTurntail = turnHead
	}

	return fakeHead.Next
}

func buildList(values []int) (head *ListNode) {
	nodes := make([]*ListNode, len(values))
	for i := 0; i < len(values); i++ {
		nodes[i] = &ListNode{
			Val:  values[i],
			Next: nil,
		}
	}
	for i := 0; i < len(values)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	return nodes[0]
}

func (sol *Solution) Title25() {
	ls := []int{1, 2, 3, 4, 5}
	k := 2
	head := CreateList(ls)
	head = reverseKGroup(head, k)
	PrintList(head)
}
