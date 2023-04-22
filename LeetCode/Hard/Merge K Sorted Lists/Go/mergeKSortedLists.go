package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func merge(list []*ListNode) *ListNode {
	if len(list) == 0 {
		return nil
	}

	n := len(list) - 1

	for n != 0 {
		i, j := 0, n
		for i < j {
			list[i] = mergeTwoLists(list[i], list[j])
			i++
			j--
			if i >= j {
				n = j
			}
		}
	}
	return list[0]
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l3 := New(-1)
	res := l3
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l3.Next = New(l2.Val)
			l2 = l2.Next
		} else {
			l3.Next = New(l1.Val)
			l1 = l1.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		for l1 != nil {
			l3.Next = New(l1.Val)
			l1 = l1.Next
			l3 = l3.Next
		}
	}

	if l2 != nil {
		for l2 != nil {
			l3.Next = New(l2.Val)
			l2 = l2.Next
			l3 = l3.Next
		}
	}
	return res.Next
}

func (l *ListNode) String() string {
	str := ""
	for l != nil {
		if l.Next == nil {
			str += strconv.Itoa(l.Val)
		} else {
			str += strconv.Itoa(l.Val) + " "
		}
		l = l.Next
	}
	return str
}

func main() {
	l1 := New(1)
	l1.Next = New(4)
	l1.Next.Next = New(5)
	l2 := New(1)
	l2.Next = New(3)
	l2.Next.Next = New(4)
	l3 := New(2)
	l3.Next = New(6)
	list := []*ListNode{l1, l2, l3}
	fmt.Println(merge(list))
}
