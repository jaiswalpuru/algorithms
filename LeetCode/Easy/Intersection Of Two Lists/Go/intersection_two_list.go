package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func intersection_two_list(a, b *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	for a != nil {
		visited[a] = true
		a = a.Next
	}

	for b != nil {
		if visited[b] {
			return b
		}
		b = b.Next
	}
	return nil
}

func main() {
	a := New(4)
	a.Next = New(1)
	a.Next.Next = New(8)
	a.Next.Next.Next = New(4)
	a.Next.Next.Next.Next = New(5)
	b := New(5)
	b.Next = New(6)
	b.Next.Next = New(1)
	b.Next.Next.Next = a.Next.Next
	fmt.Println(intersection_two_list(a, b))
}
