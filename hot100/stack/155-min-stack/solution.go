package main

import "math"

type element struct {
	value    int
	minValue int
}

type MinStack struct {
	stack []element
}

func Constructor() MinStack {
	return MinStack{stack: make([]element, 0)}
}

func (this *MinStack) Push(value int) {
	lastMin := math.MaxInt
	if len(this.stack) > 0 {
		lastMin = this.stack[len(this.stack)-1].minValue
	}

	ele := element{value: value}
	if value < lastMin {
		ele.minValue = value
	} else {
		ele.minValue = lastMin
	}

	this.stack = append(this.stack, ele)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].value
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
