package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val} }

func plus_one_linked_list(node *ListNode) *ListNode {
	c := 1
	helper(node, &c)
	if c != 0 {
		new_node := New(c)
		new_node.Next = node
		return new_node
	}
	return node
}

func helper(node *ListNode, c *int) {
	if node == nil {
		return
	}
	helper(node.Next, c)
	sum := *c + node.Val
	*c = sum / 10
	sum = sum % 10
	node.Val = sum
}

func main() {
	node := New(1)
	node.Next = New(2)
	node.Next.Next = New(3)
	fmt.Println(plus_one_linked_list(node))
}
