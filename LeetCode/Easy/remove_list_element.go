/*
Given the head of a linked list and an integer val,remove all the nodes of the linked list that has
Node.val == val, and return the new head.
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func remove_elements(node *ListNode, val int) *ListNode {
	var prev, curr *ListNode = &ListNode{}, node
	temp := prev
	for curr != nil {
		if curr.Val != val {
			prev.Next = curr
			prev = prev.Next
		}
		curr = curr.Next
	}
	prev.Next = nil
	return temp.Next
}

func main() {
	node := New(1)
	node.Next = New(2)
	node.Next.Next = New(3)
	node.Next.Next.Next = New(6)
	node.Next.Next.Next.Next = New(4)
	node.Next.Next.Next.Next.Next = New(5)
	node.Next.Next.Next.Next.Next.Next = New(6)
	node = remove_elements(node, 6)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
