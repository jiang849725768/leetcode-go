package structures

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

// 结点i的左子结点为2i+1,右子结点为2i+2
// 结点i的父节点为(i-1)/2
// move the node up
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

// move the node down
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

func PushHeap(x int, h Heap) Heap {
	h = append(h, x)
	h.up(len(h) - 1)
	return h
}

func PopHeap(h Heap) (Heap, int) {
	tail := len(h) - 1
	top := h[0]
	h.swap(0, tail)
	h = h[:tail]
	h.down(0)
	return h, top
}
