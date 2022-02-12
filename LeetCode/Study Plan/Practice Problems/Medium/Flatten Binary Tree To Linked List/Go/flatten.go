package main

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func pre_order(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}
	*arr = append(*arr, root)
	pre_order(root.Left, arr)
	pre_order(root.Right, arr)
}

func flatten(root *TreeNode) {
	nodes := []*TreeNode{}
	pre_order(root, &nodes)
	n := len(nodes)
	for i := 1; i < n; i++ {
		nodes[i-1].Left = nil
		nodes[i-1].Right = nodes[i]
	}
}

//---------------------------recursive flatten---------------------------
func flatten_recursive(root *TreeNode) {
	_flatten_recursive(root)
}

func _flatten_recursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	left := _flatten_recursive(root.Left)

	right := _flatten_recursive(root.Right)

	if left != nil {
		left.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if right == nil {
		return left
	}
	return right
}

//-----------------------------------------------------------------------

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(5)
	root.Left.Left = New(3)
	root.Left.Right = New(4)
	root.Right.Right = New(6)
	flatten_recursive(root)
}
