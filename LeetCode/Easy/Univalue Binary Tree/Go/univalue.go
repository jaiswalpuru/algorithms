package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func is_unival(root *TreeNode) bool {
	val := root.Val
	return _is_unival(root, val)
}

func _is_unival(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	return _is_unival(root.Left, val) && _is_unival(root.Right, val)
}

func main() {
	root := New(1)
	root.Left = New(1)
	root.Right = New(2)
}
