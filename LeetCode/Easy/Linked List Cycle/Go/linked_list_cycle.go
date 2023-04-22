package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func has_cycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	sp, fp := head, head

	for fp.Next != nil && fp.Next.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
		if sp == fp {
			return true
		}
	}
	return false
}

func main() {
	l := New(3)
	l.Next = New(2)
	l.Next.Next = New(0)
	l.Next.Next.Next = New(-4)
	l.Next.Next.Next.Next = l.Next
	fmt.Println(has_cycle(l))
}
