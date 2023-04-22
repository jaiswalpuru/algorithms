package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func remove_duplicates_from_an_unsorted_linked_list(root *ListNode) *ListNode {
	mp := get_cnt(root)
	temp := &ListNode{Val: -1, Next: nil}
	res := temp
	for root != nil {
		if mp[root.Val] == 1 {
			temp.Next = &ListNode{Val: root.Val, Next: nil}
			temp = temp.Next
		}
		root = root.Next
	}
	return res.Next
}

func get_cnt(root *ListNode) map[int]int {
	mp := make(map[int]int)
	for root != nil {
		mp[root.Val]++
		root = root.Next
	}
	return mp
}

func main() {
	root := New(1)
	root.Next = New(2)
	root.Next.Next = New(3)
	root.Next.Next.Next = New(2)
	fmt.Println(remove_duplicates_from_an_unsorted_linked_list(root))
}
