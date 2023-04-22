package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func remove_nodes_from_list(root *ListNode) *ListNode {
	if root.Next == nil {
		return root
	}
	st := []*ListNode{}
	for root != nil {
		for len(st) > 0 && st[len(st)-1].Val < root.Val {
			st = st[:len(st)-1]
		}
		st = append(st, root)
		root = root.Next
	}
	if len(st) == 0 {
		return nil
	}
	res := st[0]
	temp := res
	for i := 1; i < len(st); i++ {
		temp.Next = st[i]
		temp = temp.Next
	}
	return res
}

func main() {
	root := New(5)
	root.Next = New(2)
	root.Next.Next = New(13)
	root.Next.Next.Next = New(3)
	root.Next.Next.Next.Next = New(8)
	fmt.Println(remove_nodes_from_list(root))
}
