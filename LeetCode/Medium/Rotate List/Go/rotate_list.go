package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func get_length(root *ListNode) int {
	l := 0
	for root != nil {
		l++
		root = root.Next
	}
	return l
}

func rotate_list(root *ListNode, k int) *ListNode {
	n := get_length(root)
	if n == 0 {
		return root
	}
	k = k % n
	if k == 0 {
		return root
	}

	temp := n - k
	head := root
	var prev *ListNode
	for temp > 0 {
		prev = head
		head = head.Next
		temp--
	}

	h := prev.Next
	prev.Next = nil
	for head.Next != nil {
		head = head.Next
	}
	head.Next = root

	return h
}

func main() {
	root := New(1)
	root.Next = New(2)
	root.Next.Next = New(3)
	root.Next.Next.Next = New(4)
	root.Next.Next.Next.Next = New(5)
	fmt.Println(rotate_list(root, 2))
	fmt.Println(rotate_list(nil, 0))
}
