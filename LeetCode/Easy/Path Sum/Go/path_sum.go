package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type Pair struct {
	node *TreeNode
	sum  int
}

func path_sum(root *TreeNode, target int) bool {

	if root == nil {
		return false
	}
	queue := []Pair{{node: root, sum: target - root.Val}}
	for len(queue) > 0 {
		curr := queue[0].node
		remaining := queue[0].sum
		queue = queue[1:]
		if remaining == 0 && curr.Left == nil && curr.Right == nil {
			return true
		}
		if curr.Left != nil {
			queue = append(queue, Pair{node: curr.Left, sum: remaining - curr.Left.Val})
		}
		if curr.Right != nil {
			queue = append(queue, Pair{node: curr.Right, sum: remaining - curr.Right.Val})
		}
	}
	return false
}

//using recursion
func has_path_sum(root *TreeNode, target int) bool {
	return _has_path_sum(root, target)
}

func _has_path_sum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	target -= root.Val
	if root.Left == nil && root.Right == nil {
		if target == 0 {
			return true
		}
	}
	return _has_path_sum(root.Left, target) || _has_path_sum(root.Right, target)
}

func main() {
	root := New(5)
	root.Left = New(4)
	root.Right = New(8)
	root.Left.Left = New(11)
	root.Left.Left.Left = New(7)
	root.Left.Left.Right = New(2)
	root.Right.Left = New(13)
	root.Right.Right = New(4)
	root.Right.Right.Right = New(1)
	fmt.Println(has_path_sum(root, 22))
}
