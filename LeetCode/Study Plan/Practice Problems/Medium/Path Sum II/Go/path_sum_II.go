package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func path_sum(root *TreeNode, target int) [][]int {
	res := [][]int{}
	_path_sum(root, target, &res, []int{}, 0)
	return res
}

func _path_sum(root *TreeNode, target int, res *[][]int, set []int, sum int) {
	if root == nil {
		return
	}

	//root node
	if root.Left == nil && root.Right == nil {
		sum += root.Val
		if sum == target {
			set = append(set, root.Val)
			*res = append(*res, append([]int{}, set...))
		}
		return
	}

	temp := append(set, root.Val)
	_path_sum(root.Left, target, res, temp, sum+root.Val)
	_path_sum(root.Right, target, res, temp, sum+root.Val)
}

func main() {
	root := New(5)
	root.Left = New(4)
	root.Right = New(8)
	root.Left.Left = New(11)
	root.Left.Left.Left = New(7)
	root.Left.Left.Right = New(2)
	root.Right.Right = New(4)
	root.Right.Right.Right = New(1)
	root.Right.Left = New(13)
	root.Right.Right.Left = New(5)
	fmt.Println(path_sum(root, 22))
}
