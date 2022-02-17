package main

import (
	"fmt"
	"sort"
)

var mp = map[int]string{
	0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f",
	6: "g", 7: "h", 8: "i", 9: "j", 10: "k", 11: "l", 12: "m",
	13: "n", 14: "o", 15: "p", 16: "q", 17: "r", 18: "s", 19: "t",
	20: "u", 21: "v", 22: "w", 23: "x", 24: "y", 25: "z",
}

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func smallest_leave_string(root *TreeNode) string {
	res := []string{}
	_smallest_leave_string(root, &res, "")
	sort.Strings(res)
	return res[0]
}

func _smallest_leave_string(root *TreeNode, res *[]string, str string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		str = mp[root.Val] + str
		*res = append(*res, str)
		return
	}

	temp := mp[root.Val] + str
	_smallest_leave_string(root.Left, res, temp)
	_smallest_leave_string(root.Right, res, temp)
	temp = temp[:len(temp)-1]
}

func main() {
	root := New(25)
	root.Left = New(1)
	root.Right = New(3)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	root.Right.Left = New(0)
	root.Right.Right = New(2)
	fmt.Println(smallest_leave_string(root))
}
