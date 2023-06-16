package answer

import (
	"fmt"
)

type MyStack struct {
	oneList []int
	slen    int
}

func MyStackConstructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.oneList = append(this.oneList, x)
	for i := 0; i < this.slen; i++ {
		val := this.Pop()
		this.oneList = append(this.oneList, val)
		this.slen++
	}
	this.slen++
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		panic("Pop from an empty list")
	}
	val := this.oneList[0]
	this.oneList = this.oneList[1:]
	this.slen--
	return val
}

func (this *MyStack) Top() int {
	if this.Empty() {
		panic("Top from an empty list")
	}
	return this.oneList[0]
}

func (this *MyStack) Empty() bool {
	return this.slen == 0
}

func (sol *Solution) Title225() {
	obj := MyStackConstructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
