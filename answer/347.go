package answer

import (
	"container/heap"
	"fmt"
)

type KHeap [][2]int

func (h *KHeap) Len() int { return len(*h) }
func (h *KHeap) Less(i, j int) bool {
	rh := *h
	return rh[i][1] < rh[j][1]

}
func (h *KHeap) Swap(i, j int) {
	rh := *h
	rh[i], rh[j] = rh[j], rh[i]
}

func (h *KHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *KHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func topKFrequent(nums []int, k int) []int {
	kmap := make(map[int]int, 0)
	for i := range nums {
		kmap[nums[i]]++
	}

	h := &KHeap{}
	heap.Init(h)
	for key, val := range kmap {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return res
}

func (sol *Solution) Title347() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result)
}
