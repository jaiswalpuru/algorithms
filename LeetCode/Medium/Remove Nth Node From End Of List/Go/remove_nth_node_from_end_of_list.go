package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val} }

//-----------------------not in one pass--------------------------------
func remove_nth_node_from_end_of_list(root *ListNode, n int) *ListNode {
	cnt := length(root)
	cnt = cnt - n
	if cnt == 0 {
		return root.Next
	}
	temp := root
	prev := temp
	for cnt > 0 {
		prev = temp
		temp = temp.Next
		cnt--
	}
	prev.Next = temp.Next
	return root
}

func length(root *ListNode) int {
	cnt := 0
	for root != nil {
		cnt++
		root = root.Next
	}
	return cnt
}

//-----------------------not in one pass--------------------------------

//-----------------------one pass--------------------------------
func remove_nth_node_from_end_of_list_fast(root *ListNode, n int) *ListNode {
	sp, fp := root, root
	for n > 0 {
		fp = fp.Next
		n--
	}
	prev := New(-1)
	prev.Next = root
	for fp != nil {
		prev = prev.Next
		sp = sp.Next
		fp = fp.Next
	}
	prev.Next = sp.Next
	return root
}

//-----------------------one pass--------------------------------

func main() {
	root := New(1)
	root.Next = New(2)
	root.Next.Next = New(3)
	root.Next.Next.Next = New(4)
	root.Next.Next.Next.Next = New(5)
	fmt.Println(remove_nth_node_from_end_of_list_fast(root, 2))
}
