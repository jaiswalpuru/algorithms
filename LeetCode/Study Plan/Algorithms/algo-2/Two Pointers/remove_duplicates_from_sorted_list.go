/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list. Return the linked list sorted as well.
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func delete_duplicates(head *ListNode) *ListNode {
	sentinel := ListNode{Next: head}
	prev := &sentinel
	cur := sentinel.Next
	for cur != nil && cur.Next != nil {
		p := cur
		for p != nil && p.Next != nil && p.Val == p.Next.Val {
			p = p.Next
		}
		if p != cur {
			prev.Next = p.Next
			cur = p.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return sentinel.Next
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 1, Next: nil}
	head.Next.Next = &ListNode{Val: 1, Next: nil}
	head.Next.Next.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next.Next.Next = &ListNode{Val: 3, Next: nil}
	// head.Next.Next.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	// head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}
	fmt.Println(delete_duplicates(head))
}
