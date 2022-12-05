package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val} }

func middle_of_linked_list(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	sp, fp := node, node.Next
	for fp.Next != nil && fp.Next.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}
	return sp.Next
}

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(4)
	head.Next.Next.Next.Next = New(5)
	head.Next.Next.Next.Next.Next = New(6)
	fmt.Println(middle_of_linked_list(head))
}
