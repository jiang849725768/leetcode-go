package utils

// https://www.cnblogs.com/yahuian/p/11945144.html

// Heap is the smallest heap data strcture
type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) compare(i, j int) bool {
	// less
	return h[i] < h[j]
	// more
	// return h[i] > h[j]
}

// 节点i的左子结点为2i+1,右子结点为2i+2
// 节点i的父节点为(i-1)/2
// 将节点i向上移动
func (h Heap) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.compare(i, parent) {
			h.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
}

// 将节点i向下移动
func (h Heap) down(i int) {
	for i < len(h) {
		left := 2*i + 1
		right := 2*i + 2
		if left >= len(h) {
			break
		}
		if right >= len(h) {
			if h.compare(left, i) {
				h.swap(left, i)
			}
			break
		}
		if h.compare(left, right) {
			if h.compare(left, i) {
				h.swap(left, i)
				i = left
			} else {
				break
			}
		} else {
			if h.compare(right, i) {
				h.swap(right, i)
				i = right
			} else {
				break
			}
		}
	}
}

// PushHeap 向堆中插入元素
func PushHeap(x int, h Heap) Heap {
	h = append(h, x)
	h.up(len(h) - 1)
	return h
}

// PopHeap 弹出堆顶元素
func PopHeap(h Heap) (Heap, int) {
	tail := len(h) - 1
	top := h[0]
	h.swap(0, tail)
	h = h[:tail]
	h.down(0)
	return h, top
}
