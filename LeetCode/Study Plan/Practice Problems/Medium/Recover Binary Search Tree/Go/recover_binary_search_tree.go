package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func recover_binary_search_tree(root *TreeNode) {
	order := []*TreeNode{}
	inorder(root, &order)

	n := len(order)
    x, y := -1,-1
    swap := false
    for i:=0;i<n-1;i++{
        if order[i+1].Val < order[i].Val{
            x = i+1
            if !swap{
                y = i
                swap = true
            }else{
                break
            }
        }
    }
    order[x].Val,order[y].Val = order[y].Val,order[x].Val
}

func inorder(root *TreeNode, order *[]*TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, order)
	*order = append(*order, root)
	inorder(root.Right, order)
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	recover_binary_search_tree(root)
	fmt.Println(root.Val,root.Left,root.Right)
}
