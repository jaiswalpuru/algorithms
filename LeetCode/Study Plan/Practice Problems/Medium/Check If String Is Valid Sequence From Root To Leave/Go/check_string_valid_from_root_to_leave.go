package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func check_string_valid(root *TreeNode, arr []int) bool {
	return _check_string_valid(root, arr, 0)
}

func _check_string_valid(root *TreeNode, arr []int, ind int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && ind < len(arr) {
		if root.Val == arr[ind] && ind == len(arr)-1 {
			return true
		}
		return false
	}

	if ind == len(arr) {
		return false
	}
	if root.Val == arr[ind] {
		return _check_string_valid(root.Left, arr, ind+1) || _check_string_valid(root.Right, arr, ind+1)
	}
	return false
}

func main() {
	root := New(8)
	root.Left = New(3)
	root.Left.Left = New(2)
	root.Left.Right = New(1)
	root.Left.Left.Left = New(5)
	root.Left.Left.Right = New(4)
	fmt.Println(check_string_valid(root, []int{8}))
}
