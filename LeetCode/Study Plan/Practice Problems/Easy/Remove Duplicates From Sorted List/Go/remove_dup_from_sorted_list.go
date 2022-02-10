package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func remove_dup(head *ListNode) *ListNode {
	temp := head
	prev := head
	for temp != nil {
		if temp.Val == prev.Val {
			for temp != nil && prev.Val == temp.Val {
				temp = temp.Next
			}
			prev.Next = temp
			prev = prev.Next
		} else {
			prev = temp
			temp = temp.Next

		}
	}
	return head
}

func main() {
	head := New(1)
	head.Next = New(1)
	head.Next.Next = New(2)
	fmt.Println(remove_dup(head))
}
