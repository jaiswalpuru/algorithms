package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func binary_tree_cameras(root *TreeNode) int {
	cam := 0
	if _recur(root, &cam) == -1 {
        cam++
    }
	return cam
}

func _recur(root *TreeNode, cam *int) int {
    if root == nil {
        return 1 //this node doesn't require a cam
    }

    l := _recur(root.Left, cam)
    r := _recur(root.Right, cam)

    if l==-1 || r == -1 { //cam needs to be installed here
        *cam++
        return 0 //cam is installed
    }

    if l == 0 || r == 0 { //cam already installed here
        return 1 //cam installed
    }

    return -1 //cam required
}

func main() {
	root := New(0)
	root.Left = New(1)
	root.Left.Left = New(2)
	root.Left.Right = New(3)
	fmt.Println(binary_tree_cameras(root))
}
