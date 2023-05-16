package answer

import (
	"fmt"

	. "leetcode/utils"
)

func isPalindrome(head *ListNode) bool {
	length := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		length++
	}
	medium := length / 2
	var pre, cur *ListNode
	for cur = head; medium > 0; medium-- {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	if length%2 == 1 {
		cur = cur.Next
	}
	for pre != nil && cur != nil {
		if pre.Val != cur.Val {
			return false
		}
		pre = pre.Next
		cur = cur.Next
	}
	if pre != nil || cur != nil {
		return false
	}
	return true
}

func (sol *Solution) Title234() {
	head := CreateList([]int{1, 2, 3, 2, 1})
	fmt.Println(isPalindrome(head))
	head = CreateList([]int{1, 1, 1, 2, 3})
	fmt.Println(isPalindrome(head))
}
