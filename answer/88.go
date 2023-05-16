package answer

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	n1, n2 := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		if n1 < 0 {
			nums1[i] = nums2[n2]
			n2--
		} else if n2 < 0 {
			nums1[i] = nums1[n1]
			n1--
		} else if nums1[n1] > nums2[n2] {
			nums1[i] = nums1[n1]
			n1--
		} else {
			nums1[i] = nums2[n2]
			n2--
		}
	}
}

func (sol *Solution) Title88() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
