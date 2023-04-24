package answer

func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func (sol *Solution) Title704() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	println(search2(nums, target))
}
