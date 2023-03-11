package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewL(val int) *ListNode { return &ListNode{Val: val} }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func sortedListToBST(head *ListNode) *TreeNode {
	sortedList := []int{}

	for head != nil {
		sortedList = append(sortedList, head.Val)
		head = head.Next
	}

	return bst(sortedList, 0, len(sortedList)-1)
}

func bst(arr []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	root := New(arr[mid])
	root.Left = bst(arr, l, mid-1)
	root.Right = bst(arr, mid+1, r)
	return root
}

func main() {
	head := NewL(-10)
	head.Next = NewL(-3)
	head.Next.Next = NewL(0)
	head.Next.Next.Next = NewL(5)
	head.Next.Next.Next.Next = NewL(9)
	fmt.Println(sortedListToBST(head))
}
