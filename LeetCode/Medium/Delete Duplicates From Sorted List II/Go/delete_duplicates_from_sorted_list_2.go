package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func delete_duplicate_from_list(head *ListNode) *ListNode {
	temp := New(-1)
	temp.Next = head
	prev := temp
	curr := prev.Next
	for curr != nil && curr.Next != nil {
		p := curr
		for p != nil && p.Next != nil && p.Val == p.Next.Val {
			p = p.Next
		}
		if p != curr {
			prev.Next = p.Next
			curr = p.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return temp.Next
}

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(3)
	head.Next.Next.Next.Next = New(4)
	head.Next.Next.Next.Next.Next = New(4)
	head.Next.Next.Next.Next.Next.Next = New(5)
	fmt.Println(delete_duplicate_from_list(head))
}
