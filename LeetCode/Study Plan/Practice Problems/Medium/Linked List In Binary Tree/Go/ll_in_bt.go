package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New_List(val int) *ListNode { return &ListNode{Val: val, Next: nil} }
func New_Node(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func check(curr *TreeNode, list *ListNode) bool {

	if curr == nil && list != nil {
		return false
	}
	if list == nil {
		return true
	}

	if curr.Val == list.Val {
		return check(curr.Left, list.Next) || check(curr.Right, list.Next)
	}

	return false
}

func is_sub_path(list *ListNode, root *TreeNode) bool {
	q := []*TreeNode{root}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.Val == list.Val {
			if check(curr, list) {
				return true
			}
		}

		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}

	}
	return false
}

func main() {
	list := New_List(1)
	list.Next = New_List(4)
	list.Next.Next = New_List(2)
	list.Next.Next.Next = New_List(6)
	list.Next.Next.Next.Next = New_List(8)

	root := New_Node(1)
	root.Left = New_Node(4)
	root.Right = New_Node(4)
	root.Left.Right = New_Node(2)
	root.Left.Right.Left = New_Node(1)
	root.Right.Left = New_Node(2)
	root.Right.Left.Left = New_Node(6)
	root.Right.Left.Right = New_Node(8)
	root.Right.Left.Right.Left = New_Node(1)
	root.Right.Left.Right.Right = New_Node(3)

	fmt.Println(is_sub_path(list, root))
}
