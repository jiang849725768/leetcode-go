package answer

import (
	. "leetcode/utils"
)

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	oa, ob := false, false
	for pa != pb {
		if pa == nil {
			if oa {
				return nil
			}
			pa = headB
			oa = true
		} else {
			pa = pa.Next
		}
		if pb == nil {
			if ob {
				return nil
			}
			pb = headA
			ob = true
		} else {
			pb = pb.Next
		}
	}

	return pa
}

func (sol *Solution) TitleM0207() {
	headA := CreateList([]int{4, 1, 8, 4, 5})
	headB := CreateList([]int{5, 0, 1, 8, 4, 5})
	headA.Next.Next = headB.Next.Next.Next
	result := getIntersectionNode2(headA, headB)
	PrintList(result)
}
