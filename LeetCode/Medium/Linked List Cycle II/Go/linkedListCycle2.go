package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	intersectingNode := getIntersectingNode(head)
	if intersectingNode == nil {
		return nil
	}

	p1 := head
	p2 := intersectingNode

	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

func getIntersectingNode(head *ListNode) *ListNode {
	sp := head
	fp := head

	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
		if sp == fp {
			return fp
		}
	}

	return nil
}

func main() {
	head := New(3)
	head.Next = New(2)
	head.Next.Next = New(0)
	head.Next.Next.Next = New(-4)
	head.Next.Next.Next.Next = head.Next
	fmt.Println(detectCycle(head))
}
