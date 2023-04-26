package answer

import (
	"fmt"
)

type MinStack struct {
	min   int
	stack []int
}

func Constructor3() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.min = val
	}
	this.stack = append(this.stack, val-this.min)
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		panic("stack is empty")
	}
	if this.stack[len(this.stack)-1] < 0 {
		this.min -= this.stack[len(this.stack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		panic("stack is empty")
	}
	if this.stack[len(this.stack)-1] < 0 {
		return this.min
	}
	return this.stack[len(this.stack)-1] + this.min

}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		panic("stack is empty")
	}
	return this.min
}

func (sol *Solution) Title155() {
	obj := Constructor3()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
