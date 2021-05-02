package main

import (
	"fmt"
	"strconv"
)

type List struct {
	data int
	next *List
}

func NewList(data int) *List { return &List{data: data, next: nil} }

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

func addRemaining(node, temp *List, carry *int) {

	for node != nil {
		temp.data = (node.data + (*carry)) % 10
		*carry = (node.data + (*carry)) / 10
		node = node.next
		temp = temp.next
	}
}

func non_recursive_add(node1, node2 *List) *List {
	if node1 == nil {
		return node2
	} else if node2 == nil {
		return node1
	} else {
		temp := &List{}
		sum := temp
		carry := 0
		for node1 != nil && node2 != nil {
			temp.data = (node1.data + node2.data + carry) % 10
			carry = (node1.data + node2.data) / 10
			node1 = node1.next
			node2 = node2.next
			temp.next = &List{}
			temp = temp.next
		}
		if node1 != nil {
			addRemaining(node1, temp, &carry)
		}
		if node2 != nil {
			addRemaining(node2, temp, &carry)
		}
		for carry > 0 {
			temp.data = carry % 10
			carry = carry / 10
			temp = temp.next
		}
		return sum
	}
}

func main() {
	node1 := NewList(9)
	node1.next = NewList(9)
	node2 := NewList(5)
	node2.next = NewList(2)
	fmt.Println(non_recursive_add(node1, node2))
}
