package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func count_node_equal_sum_of_descendants(root *TreeNode) int {
	cnt := 0
	inorder(root, &cnt)
	return cnt
}

func inorder(root *TreeNode, cnt *int) {
	if root == nil {
		return
	}
	inorder(root.Left, cnt)
	inorder(root.Right, cnt)
	sum := 0
	if root.Left != nil {
		sum += root.Left.Val
	}
	if root.Right != nil {
		sum += root.Right.Val
	}
	if sum == root.Val {
		*cnt += 1
	}
	root.Val = sum+root.Val
}

func main() {
	root := New(24)
	root.Left = New(3)
	root.Right = New(9)
	root.Left.Left = New(2)
	root.Left.Right = New(1)
	root.Right.Left = New(6)
	root.Right.Right = New(3)
	fmt.Println(count_node_equal_sum_of_descendants(root))
}
