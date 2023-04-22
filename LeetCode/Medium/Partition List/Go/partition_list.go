package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next :nil}}

func partition_list(list *ListNode, x int) *ListNode {
	before := New(-1)
	after := New(-1)
	t1,t2 := before, after

	for list != nil {
		if list.Val < x {
			before.Next = New(list.Val)
			before = before.Next
		} else {
			after.Next = New(list.Val)
			after = after.Next
		}
		list = list.Next
	}
	fmt.Println(before)
	before.Next = t2.Next
	return t1.Next
}

func main (){
	head := New(1)
	head.Next = New(4)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(2)
	head.Next.Next.Next.Next = New(5)
	head.Next.Next.Next.Next.Next = New(2)
	fmt.Println(partition_list(head, 3))
}