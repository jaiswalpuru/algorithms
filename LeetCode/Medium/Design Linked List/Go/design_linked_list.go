package main

type Node struct {
	Val  int
	Next *Node
}

func New(val int) *Node { return &Node{Val: val, Next: nil} }

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	var list MyLinkedList
	list.size = 0
	list.head = New(0)
	return list
}

func (this *MyLinkedList) Get(index int) int {
	if index <= this.size {
		itr := this.head
		for index >= 0 {
			index--
			itr = itr.Next
		}
		if itr == nil {
			return -1
		}
		return itr.Val
	} else {
		return -1
	}
}

func (this *MyLinkedList) AddAtHead(val int) { this.AddAtIndex(0, val) }

func (this *MyLinkedList) AddAtTail(val int) { this.AddAtIndex(this.size, val) }

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	if index < 0 {
		index = 0
	}

	this.size++
	itr := this.head
	for index > 0 {
		itr = itr.Next
		index--
	}

	node := New(val)
	node.Next = itr.Next
	itr.Next = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}

	itr := this.head
	for index > 0 {
		itr = itr.Next
		index--
	}

	itr.Next = itr.Next.Next

	this.size--
}

func main() {

}
