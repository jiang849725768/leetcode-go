// https://leetcode.cn/problems/lru-cache/

package answer

import (
	"fmt"
)

// LRUNode 为双向链表节点,同时存储一个键值对
type LRUNode struct {
	key   int
	value int
	prev  *LRUNode
	next  *LRUNode
}

// LRUCache 为LRU缓存结构, 用双向链表和一个hashmap实现
type LRUCache struct {
	use  int              // 已使用的空间
	cpy  int              // 总容量
	lmap map[int]*LRUNode // map of LRUNode ptr
	head *LRUNode
	tail *LRUNode
}

// LRUConstructor 构造LRU缓存
func LRUConstructor(capacity int) LRUCache {
	return LRUCache{
		0,
		capacity,
		make(map[int]*LRUNode, capacity+1),
		nil,
		nil,
	}
}

// freshNode 将节点移动到链表头部
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

// Get 获取key对应的value, 如果不存在则返回-1
func (this *LRUCache) Get(key int) int {
	if node, ok := this.lmap[key]; ok {
		this.freshNode(key)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 指定的key存在，仅更新key
	if node, ok := this.lmap[key]; ok {
		node.value = value
		this.freshNode(key)
		return
	}

	// 添加一个新的键值对
	newNode := &LRUNode{
		key,
		value,
		nil,
		this.head,
	}
	this.lmap[key] = newNode

	// 缓存为空
	if this.use == 0 {
		this.head, this.tail = newNode, newNode
		this.use++
	} else {
		this.head.prev = newNode
		this.head = newNode
		this.use++
	}

	// 缓存已满，删除尾部节点
	for this.use > this.cpy {
		tailKey := this.tail.key
		this.tail.prev.next = nil
		this.tail = this.tail.prev
		delete(this.lmap, tailKey)
		this.use--
	}
}

func (sol *Solution) Title146() {

	obj := LRUConstructor(1)
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
