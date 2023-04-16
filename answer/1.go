package answer

import "fmt"

func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int)
	for i := range nums {
		if j, ok := nmap[target-nums[i]]; ok {
			return []int{j, i}
		}
		nmap[nums[i]] = i
	}
	return nil
}

func (sol *Solution) Title1() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
