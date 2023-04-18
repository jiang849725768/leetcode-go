package utils

import (
	"fmt"
	"testing"
)

func TestCreateTree(t *testing.T) {
	ls := []int{3, 9, 20, -1, -1, 15, 7}
	head := CreateTree(ls)
	PrintTree(head)
}

func TestCreateList(t *testing.T) {
	ls := []int{1, 2, 3, 4, 5}
	head := CreateList(ls)
	PrintList(head)
}

func TestQuickSort(t *testing.T) {
	ls := []int{3, 2, 1, 5, 4}
	QuickSort(ls)
	fmt.Println(ls)
}
