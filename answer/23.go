package answer

import (
	"leetcode/structures"
)

func mergeKLists(lists []*structures.ListNode) *structures.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	mid := len(lists) / 2
	return mergeTwoLists(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func (sol *Solution) Title23() {
	lists := []*structures.ListNode{
		structures.CreateList([]int{1, 4, 5}),
		structures.CreateList([]int{1, 3, 4}),
		structures.CreateList([]int{2, 6}),
	}
	head := mergeKLists(lists)
	structures.PrintList(head)
}
