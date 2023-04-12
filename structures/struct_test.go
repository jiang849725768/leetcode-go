package structures

import (
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
