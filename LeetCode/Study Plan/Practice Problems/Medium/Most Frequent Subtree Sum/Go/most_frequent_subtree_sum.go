package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func most_frequent_substree_sum(root *TreeNode) []int {
	res := []int{}
	hash_map := make(map[int]int)
	sum(root, &hash_map)
	mp := make(map[int][]int)
	l := 0
	for k, v := range hash_map {
		mp[v] = append(mp[v], k)
		if l <= v {
			l = v
			res = mp[v]
		}
	}

	return res
}

func sum(root *TreeNode, hash_map *map[int]int) int {
	if root == nil {
		return 0
	}
	root_sum := root.Val + sum(root.Left, hash_map) + sum(root.Right, hash_map)
	(*hash_map)[root_sum]++
	return root_sum
}

func main() {
	root := New(5)
	root.Left = New(2)
	root.Right = New(-3)
	fmt.Println(most_frequent_substree_sum(root))
}
