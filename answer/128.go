package answer

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 1
	lmap := make(map[int]struct{}, len(nums))
	for i := range nums {
		lmap[nums[i]] = struct{}{}
	}
	for num := range lmap {
		if _, ok := lmap[num-1]; !ok {
			curNum := num
			ls := 0
			ok = true
			for ok {
				curNum++
				ls++
				_, ok = lmap[curNum]
			}
			if ls > res {
				res = ls
			}
		}
	}
	return res
}

func (sol *Solution) Title128() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
