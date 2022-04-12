package main

import "fmt"

type Node struct {
	Val int
	Left,
	Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func convert_binary_search_tree_to_sorted_dll(root *Node) *Node {
	if root == nil {
		return nil
	}
	res := []*Node{}
	inorder(root, &res)
	n := len(res)
	for i := 0; i < n; i++ {
		if i == 0 {
			res[i].Left = res[n-1]
			res[i].Right = res[(i+1)%n]
		} else if i == n-1 {
			res[i].Right = res[0]
			res[i].Left = res[n-2]
		} else {
			res[i].Left = res[i-1]
			res[i].Right = res[i+1]
		}
	}
	root = res[0]
	return root
}

func inorder(root *Node, res *[]*Node) {
	if root == nil {
		return
	}

	inorder(root.Left, res)
	*res = append(*res, root)
	inorder(root.Right, res)
}

func convert_binary_search_tree_to_sorted_dll_m2(root *Node) *Node {
	if root == nil {
		return nil
	}
	var first, last *Node
	inorder_mod(root, &first, &last)
	last.Right = first
	first.Left = last
	return first
}

func inorder_mod(root *Node, first, last **Node) {
	if root == nil {
		return
	}

	inorder_mod(root.Left, first, last)
	if *last != nil {
		(*last).Right = root
		root.Left = *last
	} else {
		*first = root
	}
	*last = root
	inorder_mod(root.Right, first, last)
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(5)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	fmt.Println(convert_binary_search_tree_to_sorted_dll(root))
}
