package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func main() {
	root := New(5)
	root.Left = New(1)
	root.Right = New(5)
	root.Left.Left = New(5)
	root.Left.Right = New(5)
	root.Right.Right = New(5)
	fmt.Println(count_unival_substree(root))
}

func count_unival_substree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cnt := 0
	_count_unival_substree(root, &cnt)
	return cnt
}

func _count_unival_substree(root *TreeNode, cnt *int) bool {
	if root.Left == nil && root.Right == nil {
		*cnt += 1
		return true
	}

	is_uni := true

	if root.Left != nil {
		is_uni = _count_unival_substree(root.Left, cnt) && is_uni && root.Left.Val == root.Val
	}

	if root.Right != nil {
		is_uni = _count_unival_substree(root.Right, cnt) && is_uni && root.Right.Val == root.Val
	}

	if is_uni {
		*cnt += 1
		return true
	}
	return false
}
