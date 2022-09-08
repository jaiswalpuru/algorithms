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

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func duplicates_subtrees(root *TreeNode) []*TreeNode {
	mp := make(map[string]int)
	ans := []*TreeNode{}
	traverse(root, &mp, &ans)
	return ans
}

func traverse(root *TreeNode, mp *map[string]int, ans *[]*TreeNode) string {
	if root == nil {
		return ""
	}
	str := strconv.Itoa(root.Val) + "," + traverse(root.Left, mp, ans) + "," + traverse(root.Right, mp, ans)
	if (*mp)[str] == 1 {
		*ans = append(*ans, root)
	}
	(*mp)[str]++
	return str
}

func main() {
	root := New(1)
	root.Left = New(2)
	root.Left.Left = New(4)
	root.Right = New(3)
	root.Right.Left = New(2)
	root.Right.Left.Left = New(4)
	root.Right.Right = New(4)
	fmt.Println(duplicates_subtrees(root))
}
