package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: make([]*Node, 0)} }

func find_path(node *Node, val int, path *[]int, flag *int) {
	if node == nil {
		return
	}
	*path = append(*path, node.Val)
	if node.Val == val {
		*flag = 1
		return
	}

	for i := 0; i < len(node.Children); i++ {
		find_path(node.Children[i], val, path, flag)
		if *flag == 1 {
			return
		}
	}
	*path = (*path)[:len(*path)-1]
	return
}

func find_min_dis(root *Node, a, b int) {
	if root == nil {
		return
	}

	val := root.Val
	flag := 0
	p1, p2 := []int{}, []int{}
	find_path(root, a, &p1, &flag)

	flag = 0
	find_path(root, b, &p2, &flag)
	j := 0

	for i := 0; i < min(len(p1), len(p2)); i++ {
		if p1[i] != p2[i] {
			val = p1[i-1]
			j = i - 1
			break
		}
	}

	d1, d2 := 0, 0
	for i := j; i < len(p1); i++ {
		if p1[i] == a {
			break
		} else {
			d1 += 1
		}
	}

	for i := j; i < len(p2); i++ {
		if p2[i] == b {
			break
		} else {
			d2 += 1
		}
	}
	val = d1 + d2
	fmt.Println(val)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	root := New(1)
	root.Children = append(root.Children, New(2))
	root.Children = append(root.Children, New(3))
	root.Children[0].Children = append(root.Children[0].Children, New(4))
	root.Children[0].Children = append(root.Children[0].Children, New(5))
	root.Children[1].Children = append(root.Children[1].Children, New(6))
	root.Children[1].Children = append(root.Children[1].Children, New(7))
	root.Children[1].Children = append(root.Children[1].Children, New(8))
	find_min_dis(root, 4, 3)

}
