package utils

import (
	"math/rand"
)

// QuickSort 快速排序
func QuickSort(nums []int) {
	var qsort func(left, right int)
	qsort = func(left, right int) {
		if left >= right {
			return
		}
		// 随机节点避免基本有序情况
		pivot := rand.Intn(right-left+1) + left
		nums[left], nums[pivot] = nums[pivot], nums[left]
		i, j := left, right
		// 频繁交换避免基本一致情况
		for i < j {
			for i < j && nums[i] < nums[j] {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
			for i < j && nums[i] < nums[j] {
				i++
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				j--
			}
		}
		qsort(left, i-1)
		qsort(j+1, right)
	}
	qsort(0, len(nums)-1)

}
