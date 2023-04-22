package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//--------------------recur--------------------------------
func house_robber_three(root *TreeNode) int {
	return recur(root, false)
}

func recur(root *TreeNode, flag bool) int {
	if root == nil {
		return 0
	}
	if flag {
		return recur(root.Left, false) + recur(root.Right, false)
	} else {
		rob := root.Val + recur(root.Left, true) + recur(root.Right, true)
		not_rob := recur(root.Left, false) + recur(root.Right, false)
		return max(rob, not_rob)
	}
}

//--------------------recur--------------------------------

//--------------------memoization--------------------------------
func house_robber_three_memo(root *TreeNode) int {
	memo := make(map[*TreeNode]map[bool]int)
	return _memo(root, false, &memo)
}

func _memo(root *TreeNode, flag bool, memo *map[*TreeNode]map[bool]int) int {
	if root == nil {
		return 0
	}
	if _, ok := (*memo)[root]; ok {
		if _, k := (*memo)[root][flag]; k {
			return (*memo)[root][flag]
		}
	}
	(*memo)[root] = make(map[bool]int)
	if flag {
		res := _memo(root.Left, false, memo) + _memo(root.Right, false, memo)
		(*memo)[root][flag] = res
	} else {
		rob := root.Val + _memo(root.Left, true, memo) + _memo(root.Right, true, memo)
		dont_rob := _memo(root.Left, false, memo) + _memo(root.Right, false, memo)
		(*memo)[root][flag] = max(rob, dont_rob)
	}
	return (*memo)[root][flag]
}

//--------------------memoization--------------------------------

func main() {
	root := New(3)
	root.Left = New(2)
	root.Left.Right = New(3)
	root.Right = New(3)
	root.Right.Right = New(1)
	fmt.Println(house_robber_three_memo(root))
}
