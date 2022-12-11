package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New_L(val int) *ListNode { return &ListNode{Val: val} }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New_T(val int) *TreeNode { return &TreeNode{Val: val} }

func list_of_depth(root *TreeNode) []*ListNode {
	mp := make(map[int][]int)
	res := []*ListNode{}
	q := []*TreeNode{root}
	lvl := 0
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			curr := q[i]
			mp[lvl] = append(mp[lvl], curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		q = q[n:]
		lvl++
	}
	for _, v := range mp {
		l := New_L(v[0])
		temp := l
		for k := 1; k < len(v); k++ {
			temp.Next = New_L(v[k])
			temp = temp.Next
		}
		res = append(res, l)
	}
	return res
}

func main() {
	root := New_T(1)
	root.Left = New_T(2)
	root.Right = New_T(3)
	root.Left.Left = New_T(4)
	root.Right.Right = New_T(5)
	l := list_of_depth(root)
	for i := 0; i < len(l); i++ {
		curr := l[i]
		for curr != nil {
			fmt.Print(curr.Val, "->")
			curr = curr.Next
		}
		fmt.Println()
	}
}
