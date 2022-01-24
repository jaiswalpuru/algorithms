package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func twin_sum_max(head *ListNode) int {
	res := math.MinInt64
	_twin_sum_max(&head, head, &res)
	return res
}

func _twin_sum_max(head **ListNode, tail *ListNode, res *int) {
	if tail == nil {
		return
	}
	_twin_sum_max(head, tail.Next, res)
	*res = max(*res, tail.Val+(*head).Val)
	*head = (*head).Next
}

func main() {
	root := New(5)
	root.Next = New(4)
	root.Next.Next = New(3)
	root.Next.Next.Next = New(2)
	root.Next.Next.Next.Next = New(1)
	fmt.Println(twin_sum_max(root))
}
