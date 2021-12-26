/*
Design a max stack data structure that supports the stack operations and supports finding the stack's maximum element.

Implement the MaxStack class:

MaxStack() Initializes the stack object.
void push(int x) Pushes element x onto the stack.
int pop() Removes the element on top of the stack and returns it.
int top() Gets the element on the top of the stack without removing it.
int peekMax() Retrieves the maximum element in the stack without removing it.
int popMax() Retrieves the maximum element in the stack and removes it. If there is more than one maximum element,
only remove the top-most one.
*/

package main

type MaxStack struct {
	stack     []int
	max_stack []int
}

func Constructor() MaxStack { return MaxStack{stack: make([]int, 0)} }

func (this *MaxStack) Push(x int) {
	val := x
	if len(this.max_stack) > 0 {
		if this.max_stack[len(this.max_stack)-1] > val {
			this.max_stack = append(this.max_stack, this.max_stack[len(this.max_stack)-1])
		} else {
			this.max_stack = append(this.max_stack, val)
		}
	} else {
		this.max_stack = append(this.max_stack, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MaxStack) Pop() int {
	val := this.stack[len(this.stack)-1]
	this.max_stack = this.max_stack[:len(this.max_stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return val
}

func (this *MaxStack) Top() int { return this.stack[len(this.stack)-1] }

func (this *MaxStack) PeekMax() int { return this.max_stack[len(this.max_stack)-1] }

func (this *MaxStack) PopMax() int {
	max := this.PeekMax()
	buffer := []int{}
	for this.Top() != max {
		buffer = append(buffer, this.Pop())
	}
	this.Pop()
	for len(buffer) > 0 {
		this.Push(buffer[len(buffer)-1])
		buffer = buffer[:len(buffer)-1]
	}
	return max
}

func main() {

}
