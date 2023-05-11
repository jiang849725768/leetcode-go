package answer

import (
	. "leetcode/utils"
)

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	fast := head
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func (sol *Solution) TitleJ22() {
	head := CreateList([]int{1, 2, 3, 4, 5})
	k := 2
	res := getKthFromEnd(head, k)
	PrintList(res)
}
