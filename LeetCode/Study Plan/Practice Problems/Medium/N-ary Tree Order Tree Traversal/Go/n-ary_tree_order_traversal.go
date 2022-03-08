package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: make([]*Node, 0)} }

func traverse(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	q := []*Node{root}
	for len(q) > 0 {
		n := len(q)
		temp := []int{}
		for i := 0; i < n; i++ {
			curr := q[0]
			q = q[1:]
			temp = append(temp, curr.Val)
			if len(curr.Children) != 0 {
				q = append(q, curr.Children...)
			}
		}
		res = append(res, temp)
	}
	return res
}

func main() {
	root := New(1)
	root.Children = append(root.Children, New(3))
	root.Children = append(root.Children, New(2))
	root.Children = append(root.Children, New(4))
	root.Children[0].Children = append(root.Children[0].Children, New(5))
	root.Children[0].Children = append(root.Children[0].Children, New(6))
	fmt.Println(traverse(root))
}
