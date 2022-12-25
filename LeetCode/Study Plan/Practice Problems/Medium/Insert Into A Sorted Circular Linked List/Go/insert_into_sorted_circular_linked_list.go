package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func New(val int) *Node { return &Node{Val: val} }

func insert_into_sorted_circular_linked_list(head *Node, val int) *Node {
	root := New(val)
	root.Next = root
	if head == nil {
		return root
	}
	insert := false
	prev := head
	curr := head.Next
	for {
		if prev.Val <= val && curr.Val >= val {
			insert = true
		} else if prev.Val > curr.Val {
			if val >= prev.Val || val <= curr.Val {
				insert = true
			}
		}
		if insert {
			prev.Next = root
			root.Next = curr
			return head
		}
		prev = curr
		curr = curr.Next
		if prev == head {
			break
		}
	}
	prev.Next = root
	root.Next = curr
	return head
}

func main() {
	root := New(3)
	root.Next = New(4)
	root.Next.Next = New(1)
	root.Next.Next.Next = root
	fmt.Println(insert_into_sorted_circular_linked_list(root, 2))
}
