/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.


Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]
Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
*/

package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type MinStack struct {
	stack           []int
	min_arrangement []int
}

func Constructor() MinStack { return MinStack{stack: make([]int, 0), min_arrangement: make([]int, 0)} }

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.min_arrangement) == 0 {
		this.min_arrangement = append(this.min_arrangement, val)
	} else {
		n := len(this.min_arrangement)
		flag := false
		temp := []int{}
		for i := 0; i < n; i++ {
			if val < this.min_arrangement[i] {
				temp = append(temp, this.min_arrangement[:i]...)
				temp = append(temp, val)
				temp = append(temp, this.min_arrangement[i:]...)
				flag = true
				break
			}
		}
		if flag == false {
			this.min_arrangement = append(this.min_arrangement, val)
		} else {
			this.min_arrangement = temp
		}
	}
}

func (this *MinStack) Pop() {
	n := len(this.stack) - 1
	val := this.stack[n]
	this.stack = this.stack[:n]
	m := len(this.min_arrangement)
	temp := []int{}
	for i := 0; i < m; i++ {
		if val == this.min_arrangement[i] {
			temp = append(temp, this.min_arrangement[:i]...)
			temp = append(temp, this.min_arrangement[i+1:]...)
			break
		}
	}
	this.min_arrangement = temp
}

func (this *MinStack) Top() int { return this.stack[len(this.stack)-1] }

func (this *MinStack) GetMin() int { return this.min_arrangement[0] }

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
