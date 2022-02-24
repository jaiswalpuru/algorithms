package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func sort(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	temp := node
	var mid *ListNode
	prev, sp, fp := temp, temp, temp
	for fp != nil && fp.Next != nil {
		prev = sp
		sp = sp.Next
		fp = fp.Next.Next
	}
	mid = prev.Next
	prev.Next = nil
	left := sort(temp)
	right := sort(mid)
	return merge(left, right)
}

func merge(l, r *ListNode) *ListNode {
	if l == nil {
		return r
	}

	if r == nil {
		return l
	}

	res := New(-1)
	head := res
	for l != nil && r != nil {
		if l.Val > r.Val {
			res.Next = New(r.Val)
			r = r.Next
		} else {
			res.Next = New(l.Val)
			l = l.Next
		}
		res = res.Next
	}
	if l != nil {
		res.Next = l
	}
	if r != nil {
		res.Next = r
	}
	return head.Next
}

func main() {
	node := New(4)
	node.Next = New(2)
	node.Next.Next = New(1)
	node.Next.Next.Next = New(3)
	fmt.Println(sort(node))
}
