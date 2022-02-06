package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func get_index(arr []int, val int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func construct(preorder, postorder []int) *TreeNode {
	return _construct(preorder, postorder)
}

func _construct(preorder, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := New(preorder[0])
	if len(preorder) == 1 {
		return root
	}

	ind := get_index(postorder, preorder[1]) + 1
	root.Left = _construct(preorder[1:ind+1], postorder[:ind])
	root.Right = _construct(preorder[ind+1:], postorder[ind:])
	return root
}

func main() {
	fmt.Println(construct([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}))
}
