package answer

import (
	"leetcode/utils"
)

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
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
	lists := []*utils.ListNode{
		utils.CreateList([]int{1, 4, 5}),
		utils.CreateList([]int{1, 3, 4}),
		utils.CreateList([]int{2, 6}),
	}
	head := mergeKLists(lists)
	utils.PrintList(head)
}
