package answer

import (
	"fmt"

	. "leetcode/utils"
)

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			res := head
			for res != slow {
				res = res.Next
				slow = slow.Next
			}
			return res
		}
	}

	return nil
}

func CreateCycleList(head *ListNode, pos int) *ListNode {
	if pos == -1 || head == nil {
		return head
	}
	tmp := head
	for i := 0; i < pos; i++ {
		tmp = tmp.Next
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = tmp
	return head
}

func (sol *Solution) Title142() {
	head := CreateList([]int{3, 2, 0, -4})
	head = CreateCycleList(head, 1)

	fmt.Println(detectCycle(head).Val)
}
