/*
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should
support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack, which means only push to top, peek/pop from
top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack
using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

Example 1:
Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false

Constraints:

1 <= x <= 9
At most 100 calls will be made to push, pop, peek, and empty.
All the calls to pop and peek are valid.

Follow-up: Can you implement the queue such that each operation is amortized O(1) time complexity? In other words,
performing n operations will take overall O(n) time even if one of those operations may take longer.
*/

package main

import "fmt"

/*
This can be done using two methods :
1) Either make the enqueue operation costly:
This method ensures that the oldest entered element is always at the
top of the stack 1 so that dequeue opertaion is just pop the stack 2.

While enqueuing make the use of stack2.
i)   While the stack 1 is not empty, push everything from stack 1 to stack 2
ii)  Push x to stack 1
iii) Push everything back to stack 1

2) Or make the deque operation costly:
This method ensures that enqueue will be done in O(1).
In this method in en-queue operation new element is added at top of the stack1. In dequeue operation, if stack2 is
empty all elements are moved to stack 2 and finally the top of the stack2 is returned.
i) Enqueue(q,x)
1) Push x to stack 1(assuming the size of stack is unlimited)
ii)Dequeue
i)  If both stacks are empty then error
ii) If stack 2 is empty, While stack 1 is not empty, push everything from stack 1 to stack 2
iii)Pop the element from stack2 and return.
*/

type MyQueue struct {
	s1, s2 []int
}

func Constructor() MyQueue { return MyQueue{s1: make([]int, 0), s2: make([]int, 0)} }

func (this *MyQueue) Pop() int {
	n, m := len(this.s1), len(this.s2)
	if n == 0 && m == 0 {
		return -1
	}

	if m == 0 {
		for i := n - 1; i >= 0; i-- {
			this.s2 = append(this.s2, this.s1[i])
			this.s1 = this.s1[:i]
		}
	}
	val := this.s2[0]
	this.s2 = this.s2[1:]
	return val
}

func (this *MyQueue) Push(x int) {
	n := len(this.s1) - 1
	for n >= 0 {
		this.s2 = append(this.s2, this.s1[n])
		this.s1 = this.s1[:n]
		n--
	}
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Peek() int {
	if len(this.s2) != 0 {
		return this.s2[0]
	}
	return this.s1[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.s1) == 0 && len(this.s2) == 0 {
		return true
	}
	return false
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
}
