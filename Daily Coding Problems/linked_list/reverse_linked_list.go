package main

import (
	"fmt"
	"strconv"
)

type List struct {
	data interface{}
	next *List
}

func NewList(data interface{}) *List { return &List{data: data, next: nil} }

func non_recursive_reverse(head *List) *List {
	var prev, current *List
	for current = head; current != nil; {
		current.next, prev, current = prev, current, current.next
	}
	return prev
}

func reverse(root *List) *List {
	if root == nil {
		return nil
	}
	temp := reverse(root.next)
	if temp == nil {
		return root
	}
	root.next.next = root
	return temp
}

func Reverse(head *List) *List {
	if head == nil {
		return nil
	}
	h := reverse(head)
	head.next = nil
	return h
}

func (root *List) String() string {
	head := root
	str := ""
	for head != nil {
		if head.next == nil {
			str += strconv.Itoa(head.data.(int))
		} else {
			str += strconv.Itoa(head.data.(int)) + " -> "
		}
		head = head.next
	}
	return fmt.Sprintf("%v", str)
}
func main() {
	l := NewList(12)
	l.next = NewList(14)
	l.next.next = NewList(16)
	l.next.next.next = NewList(18)
	fmt.Println("Before reverse : ", l)
	fmt.Println("After reverse : ", Reverse(l))
}
