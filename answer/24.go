package answer

import (
	. "leetcode/utils"
)

func swapPairs(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	cur := fakeHead
	for {
		fir := cur.Next
		if fir == nil {
			break
		}
		sec := fir.Next
		if sec == nil {
			break
		}
		cur.Next = sec
		fir.Next = sec.Next
		sec.Next = fir
		cur = fir
	}
	return fakeHead.Next
}

func (sol *Solution) Title24() {
	head := CreateList([]int{1, 2, 3, 4})
	result := swapPairs(head)
	PrintList(result)
	head = CreateList([]int{1, 2, 3, 4, 5})
	result = swapPairs(head)
	PrintList(result)
}
