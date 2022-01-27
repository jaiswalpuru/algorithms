package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func merge(l1, l2 *ListNode) *ListNode {
	res := _merge(l1, l2)
	return res
}

func _merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val > l2.Val {
		res = New(l2.Val)
		res.Next = _merge(l1, l2.Next)
	} else {
		res = New(l1.Val)
		res.Next = _merge(l1.Next, l2)
	}
	return res
}

func main() {
	l1 := New(1)
	l1.Next = New(2)
	l1.Next.Next = New(4)
	l2 := New(1)
	l2.Next = New(3)
	l2.Next.Next = New(4)
	fmt.Println(merge(l1, l2))
}
