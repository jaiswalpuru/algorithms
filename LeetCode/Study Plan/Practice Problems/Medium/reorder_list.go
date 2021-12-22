/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
*/

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func New(val int) *Node { return &Node{Val: val, Next: nil} }

func reorder_list(node *Node) *Node {
	//find the middle node
	sp, fp := node, node
	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}

	//reverse the second half
	var prev, curr *Node = nil, sp
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	//merge the two list
	f, s := node, prev
	for s.Next != nil {
		temp := f.Next
		f.Next = s
		f = temp

		temp = s.Next
		s.Next = f
		s = temp
	}

	return node
}

func main() {
	node := New(1)
	node.Next = New(2)
	node.Next.Next = New(3)
	node.Next.Next.Next = New(4)
	t := reorder_list(node)
	fmt.Println(t, t.Next)

}
