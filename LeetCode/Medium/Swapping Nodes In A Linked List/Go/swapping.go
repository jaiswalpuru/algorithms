package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func list_len(head *ListNode) int {
	n := 0
	for head != nil {
		head = head.Next
		n++
	}
	return n
}

func swap_node(head *ListNode, k int) *ListNode {
	n := list_len(head)

	start, end := head, head
	r := n - k + 1
	for k > 1 {
		start = start.Next
		k--
	}

	for r > 1 {
		end = end.Next
		r--
	}
	start.Val, end.Val = end.Val, start.Val
	return head
}

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(4)
	head.Next.Next.Next.Next = New(5)
	fmt.Println(swap_node(head, 2).Next)
}
