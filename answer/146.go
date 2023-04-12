// https://leetcode.cn/problems/lru-cache/

package answer

import (
	"fmt"
)

// LRUNode form a two-way linked list
type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

type LRUCache struct {
	use  int
	cpy  int
	lmap map[int]*LRUNode // map of LRUNode ptr
	head *LRUNode
	tail *LRUNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		0,
		capacity,
		make(map[int]*LRUNode, capacity+1),
		nil,
		nil,
	}
}

func (this *LRUCache) freshNode(key int) {
	if node, ok := this.lmap[key]; ok {
		if node == this.head {
			return
		}
		if node == this.tail {
			this.tail = this.tail.prev
			node.prev.next = nil
		} else {
			node.prev.next, node.next.prev = node.next, node.prev
		}
		node.next, node.prev = this.head, nil
		this.head.prev = node
		this.head = node
	} else {
		panic("LRU fresh unexisted key")
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.lmap[key]; ok {
		this.freshNode(key)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// key existed
	if node, ok := this.lmap[key]; ok {
		node.value = value
		this.freshNode(key)
		return
	}

	// add new key-value
	newNode := &LRUNode{
		key,
		value,
		nil,
		this.head,
	}
	this.lmap[key] = newNode

	if this.use == 0 {
		this.head, this.tail = newNode, newNode
		this.use++
	} else {
		this.head.prev = newNode
		this.head = newNode
		this.use++
	}

	for this.use > this.cpy {
		tailKey := this.tail.key
		this.tail.prev.next = nil
		this.tail = this.tail.prev
		delete(this.lmap, tailKey)
		this.use--
	}
}

func (sol *Solution) Title146() {

	obj := Constructor(1)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

}
