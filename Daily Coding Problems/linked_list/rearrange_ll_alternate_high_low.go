package main

import (
	"fmt"
	"strconv"
)

type List struct {
	data int
	next *List
}

func (root *List) String() string {
	head := root
	str := ""
	for head != nil {
		if head.next == nil {
			str += strconv.Itoa(head.data)
		} else {
			str += strconv.Itoa(head.data) + " -> "
		}
		head = head.next
	}
	return fmt.Sprintf("%v", str)
}

func NewList(data int) *List { return &List{data: data, next: nil} }

func alternate(head *List) *List {
	even := true
	temp := head
	for temp.next != nil {
		if temp.data > temp.next.data && even {
			temp.data, temp.next.data = temp.next.data, temp.data
		} else if temp.data < temp.next.data && !even {
			temp.data, temp.next.data = temp.next.data, temp.data
		}
		even = !even
		temp = temp.next
	}
	return head
}

func alternate_without_parity(head *List) *List {
	prev := head
	curr := prev.next
	for curr != nil {
		if prev.data > curr.data {
			prev.data, curr.data = curr.data, prev.data
		}

		if curr.next == nil {
			break
		}

		if curr.next.data > curr.data {
			curr.next.data, curr.data = curr.data, curr.next.data
		}

		prev = curr.next
		curr = curr.next.next
	}
	return head
}

func main() {
	head := NewList(1)
	head.next = NewList(2)
	head.next.next = NewList(3)
	head.next.next.next = NewList(4)
	head.next.next.next.next = NewList(5)
	fmt.Println("Before alternating : ", head)
	fmt.Println("After alternating : ", alternate(head))
}
