package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val} }

func odd_even_linked_list(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	i := 0
	n := 0
	t, temp := node, node
	for temp.Next != nil {
		temp = temp.Next
		n++
	}
	n += 1
	prev := t
	for n != 0 {
		if i%2 == 1 {
			temp.Next = New(t.Val)
			temp = temp.Next
			prev.Next = t.Next
			t = prev.Next
		} else {
			prev = t
			t = t.Next
		}
		n--
		i++
	}
	return node
}

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(4)
	head.Next.Next.Next.Next = New(5)
	fmt.Println(odd_even_linked_list(head))
}
