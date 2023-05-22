package answer

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if nums[i]+nums[j]+nums[len(nums)-1]+nums[len(nums)-2] < target {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			l, r := j+1, len(nums)-1
			for l < r {
				if nums[i]+nums[j]+nums[l]+nums[r] == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if nums[i]+nums[j]+nums[l]+nums[r] < target {
					l++
				} else {
					r--
				}
			}
			for j < len(nums)-3 && nums[j] == nums[j+1] {
				j++
			}
		}
	}

	return res
}

func (sol *Solution) Title18() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(nums, target)
	fmt.Println(result)
}
