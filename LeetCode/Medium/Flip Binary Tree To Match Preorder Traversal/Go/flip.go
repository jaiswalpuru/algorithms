package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func flip_binary_tree(root *TreeNode, preorder []int) []int {
	flipped := []int{}
	ind := 0
	dfs(root, preorder, &ind, &flipped)
	if len(flipped) > 0 && flipped[0] == -1 {
		return []int{-1}
	}
	return flipped
}

func dfs(root *TreeNode, preorder []int, ind *int, flipped *[]int) {
	if root != nil {

		if root.Val != preorder[*ind] {
			*flipped = []int{-1}
			return
		}
		*ind++
		if *ind < len(preorder) && root.Left != nil && root.Left.Val != preorder[*ind] {
			*flipped = append(*flipped, root.Val)
			dfs(root.Right, preorder, ind, flipped)
			dfs(root.Left, preorder, ind, flipped)
		} else {
			dfs(root.Left, preorder, ind, flipped)
			dfs(root.Right, preorder, ind, flipped)
		}
	}
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	fmt.Println(flip_binary_tree(root, []int{1, 3, 2}))
}
