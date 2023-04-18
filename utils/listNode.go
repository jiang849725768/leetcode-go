package utils

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateList create a linked list with int slice
func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	node := head
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{Val: nums[i]}
		node = node.Next
	}
	return head
}

// PrintList prinr value of list in order
func PrintList(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Printf("%d ", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Printf("\n")
}
