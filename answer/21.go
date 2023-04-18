package answer

import (
	"leetcode/utils"
)

func mergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	fakeHead := &utils.ListNode{
		Next: nil,
	}
	tmp := fakeHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}
	if list1 != nil {
		tmp.Next = list1
	}
	if list2 != nil {
		tmp.Next = list2
	}

	return fakeHead.Next
}
