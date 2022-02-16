package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func swap_nodes_in_pair(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	remaining := head.Next.Next
	new_node := head.Next
	head.Next.Next = head
	head.Next = swap_nodes_in_pair(remaining)
	return new_node
}

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(4)
	fmt.Println(swap_nodes_in_pair(head))
}
