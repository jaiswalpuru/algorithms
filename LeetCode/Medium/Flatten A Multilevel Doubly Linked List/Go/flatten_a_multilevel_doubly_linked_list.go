package main

import "fmt"

type Node struct {
	Val               int
	Prev, Next, Child *Node
}

func New(val int) *Node { return &Node{Val: val} }

func flatten_a_multilevel_doubly_linked_list(root *Node) *Node {
	if root == nil {
		return nil
	}
	st := []*Node{}
	dfs(root, &st)
	head := New(st[0].Val)
	t := head
	for i := 1; i < len(st); i++ {
		curr := st[i]
		curr.Child = nil
		t.Next = curr
		curr.Prev = t
		t = t.Next
	}
	return head
}

func dfs(root *Node, st *[]*Node) {
	if root == nil {
		return
	}
	*st = append(*st, root)
	if root.Child != nil {
		dfs(root.Child, st)
	}
	dfs(root.Next, st)
}

func main() {
	root := New(1)
	root.Next = New(2)
	root.Next.Prev = root
	root.Next.Next = New(3)
	root.Next.Next.Prev = root.Next
	root.Next.Next.Next = New(4)
	root.Next.Next.Next.Prev = root.Next.Next
	root.Next.Next.Next.Next = New(5)
	root.Next.Next.Next.Next.Prev = root.Next.Next.Next
	root.Next.Next.Next.Next.Next = New(6)
	root.Next.Next.Next.Next.Next.Prev = root.Next.Next.Next.Next
	c1 := New(7)
	c1.Next = New(8)
	c1.Next.Prev = c1
	c1.Next.Next = New(9)
	c1.Next.Next.Prev = c1.Next
	c1.Next.Next.Next = New(10)
	c1.Next.Next.Next.Prev = c1.Next.Next
	root.Next.Next.Child = c1
	c2 := New(11)
	c2.Next = New(12)
	c2.Next.Prev = c2
	c1.Next.Child = c2
	fmt.Println(flatten_a_multilevel_doubly_linked_list(root))
}
