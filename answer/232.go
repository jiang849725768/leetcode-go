package answer

import (
	"fmt"
)

type MyQueue struct {
	instack  []int
	outstack []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.instack = append(this.instack, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		panic("The queue is empty")
	}
	if len(this.outstack) == 0 {
		for len(this.instack) > 0 {
			this.outstack = append(this.outstack, this.instack[len(this.instack)-1])
			this.instack = this.instack[:len(this.instack)-1]
		}
	}
	res := this.outstack[len(this.outstack)-1]
	this.outstack = this.outstack[:len(this.outstack)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		panic("The queue is empty")
	}
	if len(this.outstack) != 0 {
		return this.outstack[len(this.outstack)-1]
	} else {
		return this.instack[0]
	}
}

func (this *MyQueue) Empty() bool {
	return len(this.instack) == 0 && len(this.outstack) == 0
}

func (sol *Solution) Title232() {
	obj := MyQueueConstructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Push(5)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
}
