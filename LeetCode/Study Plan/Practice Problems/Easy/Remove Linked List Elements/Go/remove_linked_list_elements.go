package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val} }

func remove_linked_list_elements(head *ListNode, val int) *ListNode {
	prev := New(-1)
	temp := prev
	prev.Next = head
	for head != nil {
		if head.Val == val {
			temp.Next = head.Next
			head = head.Next
		} else {
			temp = head
			head = head.Next
		}
	}
	return prev.Next
}

func main() {
	node := New(1)
	node.Next = New(2)
	node.Next.Next = New(6)
	node.Next.Next.Next = New(3)
	node.Next.Next.Next.Next = New(4)
	node.Next.Next.Next.Next.Next = New(5)
	node.Next.Next.Next.Next.Next.Next = New(6)
	fmt.Println(remove_linked_list_elements(node, 6))
}
