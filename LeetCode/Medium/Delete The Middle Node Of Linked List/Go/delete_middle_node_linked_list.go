package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func delete_middle_node_linked_list(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	sp, fp, prev := head, head, head

	for fp != nil && fp.Next != nil {
		prev = sp
		sp = sp.Next
		fp = fp.Next.Next
	}
	if sp.Next != nil {
		prev.Next = sp.Next
	} else {
		prev.Next = nil
	}
	return head
}

func main() {
	node := New(1)
	node.Next = New(3)
	node.Next.Next = New(4)
	node.Next.Next.Next = New(7)
	node.Next.Next.Next.Next = New(1)
	node.Next.Next.Next.Next.Next = New(2)
	node.Next.Next.Next.Next.Next.Next = New(6)
	fmt.Println(delete_middle_node_linked_list(node))
}
