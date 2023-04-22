package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func binary_tree_paths(root *TreeNode) []string {
	res := [][]int{}
	paths(root, &res, []int{})
	ans := []string{}
	n := len(res)
	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < len(res[i]); j++ {
			if j == len(res[i])-1 {
				str += strconv.Itoa(res[i][j])
			} else {
				str += strconv.Itoa(res[i][j]) + "->"
			}
		}
		ans = append(ans, str)
	}
	return ans
}

func paths(root *TreeNode, res *[][]int, str []int) {
	if root == nil {
		return
	}

	temp := append(str, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, append([]int{}, temp...))
	} else {
		paths(root.Left, res, temp)
		paths(root.Right, res, temp)
	}
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Right = New(5)
	root.Right = New(3)
	fmt.Println(binary_tree_paths(root))
}
