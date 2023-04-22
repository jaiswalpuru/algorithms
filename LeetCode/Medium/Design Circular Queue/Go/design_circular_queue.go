package main

//------------------using array as a ring list---------------------
type MyCircularQueue struct {
	q        []int
	head     int
	count    int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{q: make([]int, k), head: 0, count: 0, capacity: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.count == this.capacity {
		return false
	}
	this.q[(this.head+this.count)%this.capacity] = value
	this.count++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.count == 0 {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	this.count -= 1
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}
	return this.q[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.count == 0 {
		return -1
	}
	tail := (this.head + this.count - 1) % this.capacity
	return this.q[tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.capacity
}

//------------------using array as a ring list---------------------

//------------------using linked list---------------------
type Node struct {
	Val  int
	next *Node
}

func New(val int) *Node { return &Node{Val: val, next: nil} }

type MyCircularQueue struct {
	head *Node
	k, l int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{head: New(-1), k: k, l: 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.k == this.l {
		return false
	}
	this.l++
	if this.head.Val == -1 {
		this.head.Val = value
		this.head.next = this.head
	} else {
		temp := this.head
		for temp.next != this.head {
			temp = temp.next
		}
		temp.next = New(value)
		temp.next.next = this.head
	}
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.head.Val == -1 {
		return false
	}
	this.l--
	if this.head.next == this.head {
		this.head = New(-1)
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

func (this *MyCircularQueue) Front() int {
	if this.head.Val == -1 {
		return -1
	}
	return this.head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.head.Val == -1 {
		return -1
	}
	temp := this.head
	for temp.next != this.head {
		temp = temp.next
	}
	return temp.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head.Val == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return this.l == this.k
}

//------------------using linked list---------------------
