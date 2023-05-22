package answer

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}
	m := make(map[int]bool)
	for i := range nums2 {
		m[nums2[i]] = true
	}
	res := make([]int, 0)
	for i := range nums1 {
		if m[nums1[i]] {
			res = append(res, nums1[i])
			m[nums1[i]] = false
		}
	}
	return res
}

func (sol *Solution) Title349() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersection(nums1, nums2)
	fmt.Println(result)
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	result = intersection(nums1, nums2)
	fmt.Println(result)
}
