package answer

import (
	. "leetcode/utils"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	tmp1, tmp2 := l1, l2
	for {
		tmp1.Val += tmp2.Val + carry
		carry = tmp1.Val / 10
		tmp1.Val %= 10
		if tmp1.Next == nil || tmp2.Next == nil {
			break
		}
		tmp1, tmp2 = tmp1.Next, tmp2.Next
	}
	if tmp1.Next == nil {
		tmp1.Next = tmp2.Next
	}
	for tmp1.Next != nil {
		tmp1 = tmp1.Next
		tmp1.Val += carry
		carry = tmp1.Val / 10
		tmp1.Val %= 10
	}
	if carry != 0 {
		tmp1.Next = &ListNode{
			Val: carry,
		}
	}

	return l1
}

func (sol *Solution) Title2() {
	l1 := CreateList([]int{2, 4, 5})
	l2 := CreateList([]int{5, 6, 4})
	PrintList(l1)
	PrintList(l2)
	PrintList(addTwoNumbers(l1, l2))
}
