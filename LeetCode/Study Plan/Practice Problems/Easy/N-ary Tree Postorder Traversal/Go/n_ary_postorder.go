package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: make([]*Node, 0)} }

func postorder(root *Node) []int {
	res := []int{}
	_postorder(root, &res)
	return res
}

func _postorder(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for i := 0; i < len(root.Children); i++ {
		_postorder(root.Children[i], res)
	}
	*res = append(*res, root.Val)
}

func main() {
	root := New(1)
	root.Children = append(root.Children, New(3))
	root.Children = append(root.Children, New(2))
	root.Children = append(root.Children, New(4))
	root.Children[0].Children = append(root.Children[0].Children, New(5))
	root.Children[0].Children = append(root.Children[0].Children, New(6))
	fmt.Println(postorder(root))
}
