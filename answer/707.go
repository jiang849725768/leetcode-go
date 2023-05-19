package answer

import (
	"fmt"
)

type MyLinkedListNode struct {
	val  int
	next *MyLinkedListNode
	prev *MyLinkedListNode
}

type MyLinkedList struct {
	head   *MyLinkedListNode
	tail   *MyLinkedListNode
	length int
}

func MyLinkedListConstructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
	tmp := this.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}
	return tmp.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedListNode{val: val}
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &MyLinkedListNode{val: val}
	if this.tail == nil {
		this.head = node
		this.tail = node
	} else {
		node.prev = this.tail
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.length {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.length {
		this.AddAtTail(val)
		return
	}
	tmp := this.head
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}
	node := &MyLinkedListNode{val: val}
	node.prev = tmp.prev
	node.next = tmp
	tmp.prev.next = node
	tmp.prev = node
	this.length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}
	if index == 0 {
		if this.length == 1 {
			this.head = nil
			this.tail = nil
		} else {
			this.head = this.head.next
			this.head.prev = nil
		}
	} else if index == this.length-1 {
		this.tail = this.tail.prev
		this.tail.next = nil
	} else {
		tmp := this.head
		for i := 0; i < index; i++ {
			tmp = tmp.next
		}
		tmp.prev.next = tmp.next
		tmp.next.prev = tmp.prev
	}
	this.length--
}

func (sol *Solution) Title707() {
	obj := MyLinkedListConstructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	fmt.Println(obj.Get(1))
	obj.DeleteAtIndex(1)
	fmt.Println(obj.Get(1))
}
