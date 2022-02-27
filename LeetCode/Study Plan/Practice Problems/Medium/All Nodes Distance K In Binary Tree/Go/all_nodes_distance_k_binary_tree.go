package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func all_node_distance_k_binary_tree(root *TreeNode, target *TreeNode, k int) []int {
	list := []int{}
	dfs(root, &list, target, k)
	return list
}

func dfs(root *TreeNode, list *[]int, target *TreeNode, k int) int {
	if root == nil {
		return -1
	} else if root == target {
		sub_tree_add(root, list, k, 0)
		return 1
	} else {
		l := dfs(root.Left, list, target, k)
		r := dfs(root.Right, list, target, k)
		if l != -1 {
			if l == k {
				*list = append(*list, root.Val)
			}
			sub_tree_add(root.Right, list, k, l+1)
			return l + 1
		} else if r != -1 {
			if r == k {
				*list = append(*list, root.Val)
			}
			sub_tree_add(root.Left, list, k, r+1)
			return r + 1
		} else {
			return -1
		}
	}
}

func sub_tree_add(root *TreeNode, list *[]int, k int, dist int) {
	if root == nil {
		return
	}
	if dist == k {
		*list = append(*list, root.Val)
	} else {
		sub_tree_add(root.Left, list, k, dist+1)
		sub_tree_add(root.Right, list, k, dist+1)
	}
}

func main() {
	root := New(0)
	root.Left = New(1)
	root.Left.Left = New(3)
	root.Left.Right = New(2)
	fmt.Println(all_node_distance_k_binary_tree(root, root.Left.Right, 1))
}
