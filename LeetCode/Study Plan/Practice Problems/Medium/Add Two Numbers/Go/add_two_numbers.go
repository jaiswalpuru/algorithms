package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func add_two_number(l1, l2 *ListNode) *ListNode {
	return _add(l1, l2, 0)
}

func _add(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 == nil && carry != 0 {
		return New(carry)
	}
	val := carry
	var new_node *ListNode

	if l1 != nil && l2 != nil {
		val += l1.Val + l2.Val
		l1 = l1.Next
		l2 = l2.Next
	} else if l2 == nil {
		val += l1.Val
		l1 = l1.Next
	} else {
		val += l2.Val
		l2 = l2.Next
	}
	new_node = New(val % 10)
	carry = val / 10
	new_node.Next = _add(l1, l2, carry)
	return new_node
}

func main() {
	l1 := New(2)
	l1.Next = New(4)
	l1.Next.Next = New(3)
	l2 := New(5)
	l2.Next = New(6)
	l2.Next.Next = New(4)
	fmt.Println(add_two_number(l1, l2))
}
