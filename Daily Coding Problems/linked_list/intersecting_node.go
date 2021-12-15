package main

import (
	"fmt"
	"strconv"
)

type List struct {
	data int
	next *List
}

func Length(head *List) int {
	if head == nil {
		return 0
	}
	return 1 + Length(head.next)
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

func find_intersect(list1, list2 *List) *List {
	l1, l2 := list1, list2
	m, n := Length(list1), Length(list2)
	if m > n {
		for i := 0; i < (m - n); i++ {
			l1 = l1.next
		}
	} else {
		for i := 0; i < (n - m); i++ {
			l2 = l2.next
		}
	}

	for l1 != l2 {
		l1 = l1.next
		l2 = l2.next
	}
	return l1
}

func main() {
	A := NewList(3)
	A.next = NewList(7)
	A.next.next = NewList(8)
	A.next.next.next = NewList(10)
	B := NewList(99)
	B.next = NewList(1)
	B.next.next = A.next.next
	fmt.Println(find_intersect(A, B).data)
}
