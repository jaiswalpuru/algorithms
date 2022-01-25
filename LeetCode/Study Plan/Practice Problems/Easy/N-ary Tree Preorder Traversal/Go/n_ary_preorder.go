package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: make([]*Node, 0)} }

func preorder(root *Node) []int {
	res := []int{}
	_preorder(root, &res)
	return res
}

func _preorder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for i := 0; i < len(root.Children); i++ {
		_preorder(root.Children[i], res)
	}
}

func main() {
	root := New(1)
	root.Children = append(root.Children, New(3))
	root.Children = append(root.Children, New(2))
	root.Children = append(root.Children, New(4))
	root.Children[0].Children = append(root.Children[0].Children, New(5))
	root.Children[0].Children = append(root.Children[0].Children, New(6))
	fmt.Println(preorder(root))
}
