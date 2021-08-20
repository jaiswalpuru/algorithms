/*
Design your implementation of the circular queue. The circular queue is a linear data structure in
which the operations are performed based on FIFO (First In First Out) principle and the last position
is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue.
In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space
in front of the queue. But using the circular queue, we can use the space to store new values.

Implementation the MyCircularQueue class:

MyCircularQueue(k) Initializes the object with the size of the queue to be k.
int Front() Gets the front item from the queue. If the queue is empty, return -1.
int Rear() Gets the last item from the queue. If the queue is empty, return -1.
boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
boolean isEmpty() Checks whether the circular queue is empty or not.
boolean isFull() Checks whether the circular queue is full or not.
You must solve the problem without using the built-in queue data structure in your programming language.

Example 1:
Input
["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
Output
[null, true, true, true, false, 3, true, true, true, 4]

Explanation
MyCircularQueue myCircularQueue = new MyCircularQueue(3);
myCircularQueue.enQueue(1); // return True
myCircularQueue.enQueue(2); // return True
myCircularQueue.enQueue(3); // return True
myCircularQueue.enQueue(4); // return False
myCircularQueue.Rear();     // return 3
myCircularQueue.isFull();   // return True
myCircularQueue.deQueue();  // return True
myCircularQueue.enQueue(4); // return True
myCircularQueue.Rear();     // return 4
*/

package main

import "fmt"

type Val struct {
	val  int
	next *Val
}

type MyCircularQueue struct {
	head *Val
	k    int
	l    int
}

func Constructor(k int) MyCircularQueue { return MyCircularQueue{head: &Val{val: -1}, k: k, l: 0} }

func (this *MyCircularQueue) EnQueue(val int) bool {
	if this.l == this.k {
		return false
	}
	this.l++
	if this.head.val == -1 {
		this.head.val = val
		this.head.next = this.head
	} else {
		temp := this.head
		for temp.next != this.head {
			temp = temp.next
		}
		new_val := &Val{val: val, next: nil}
		temp.next = new_val
		temp.next.next = this.head
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.head.val == -1 {
		return false
	}

	this.l--
	if this.head.next == this.head {
		this.head = &Val{val: -1}
		return true
	}
	temp := this.head
	for temp.next != this.head {
		temp = temp.next
	}
	temp.next = this.head.next
	this.head = temp.next
	return true
}

func (this *MyCircularQueue) Front() int { return this.head.val }

func (this *MyCircularQueue) Rear() int {
	if this.head.val == -1 {
		return -1
	}
	temp := this.head
	for temp.next != this.head {
		temp = temp.next
	}
	return temp.val
}

func (this *MyCircularQueue) IsEmpty() bool { return this.head.val == -1 }

func (this *MyCircularQueue) IsFull() bool { return this.l == this.k }

func main() {
	obj := Constructor(3)
	fmt.Println(obj.EnQueue(7))
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Front())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Front())
	fmt.Println(obj.Rear())
	fmt.Println(obj.EnQueue(0))
	fmt.Println(obj.IsFull())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Rear())
	fmt.Println(obj.EnQueue(3))
}
