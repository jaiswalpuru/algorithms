package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func create_binary_tree(d [][]int) *TreeNode {
	mp := make(map[int]*TreeNode)
	n := len(d)
	for i := 0; i < n; i++ {
		if _, ok := mp[d[i][0]]; !ok {
			node := New(d[i][0])
			mp[d[i][0]] = node
			if d[i][2] == 1 {
				if _, ok := mp[d[i][1]]; ok {
					node.Left = mp[d[i][1]]
				} else {
					node.Left = New(d[i][1])
					mp[d[i][1]] = node.Left
				}
				mp[d[i][0]].Left = node.Left
			} else {
				if _, ok := mp[d[i][1]]; ok {
					node.Right = mp[d[i][1]]
				} else {
					node.Right = New(d[i][1])
					mp[d[i][1]] = node.Right
				}
				mp[d[i][0]].Right = node.Right
			}
		} else {
			node := mp[d[i][0]]
			if d[i][2] == 1 {
				if _, ok := mp[d[i][1]]; ok {
					node.Left = mp[d[i][1]]
				} else {
					node.Left = New(d[i][1])
					mp[d[i][1]] = node.Left
				}
				mp[d[i][0]].Left = node.Left
			} else {
				if _, ok := mp[d[i][1]]; ok {
					node.Right = mp[d[i][1]]
				} else {
					node.Right = New(d[i][1])
					mp[d[i][1]] = node.Right
				}
				mp[d[i][0]].Right = node.Right
			}
		}
	}
	parent := make(map[int]bool)
	var root *TreeNode
	for i := 0; i < n; i++ {
		if _, ok := parent[d[i][1]]; !ok {
			parent[d[i][1]] = true
		}
	}

	for i := 0; i < n; i++ {
		if _, ok := parent[d[i][0]]; !ok {
			root = mp[d[i][0]]
			break
		}
	}

	return root
}

func main() {
	fmt.Println(create_binary_tree([][]int{
		{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1},
	}))
}
