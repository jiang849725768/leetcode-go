package answer

import (
	"fmt"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	nmap := make(map[int]int, len(nums1)*len(nums2))
	for i := range nums1 {
		for j := range nums2 {
			nmap[nums1[i]+nums2[j]]++
		}
	}
	res := 0
	for i := range nums3 {
		for j := range nums4 {
			if v, ok := nmap[-nums3[i]-nums4[j]]; ok {
				res += v
			}
		}
	}

	return res
}

func (sol *Solution) Title454() {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	result := fourSumCount(nums1, nums2, nums3, nums4)
	fmt.Println(result)
}
