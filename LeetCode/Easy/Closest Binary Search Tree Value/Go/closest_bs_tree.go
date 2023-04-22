package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func closest_val(root *TreeNode, val float64) int {
	if root == nil {
		return 0
	}

	ans := math.MaxFloat64
	v := -1
	_closest_val(root, val, &ans, &v)
	return v
}

func _closest_val(root *TreeNode, val float64, ans *float64, v *int) {
	if root == nil {
		return
	}
	_closest_val(root.Left, val, ans, v)
	if float64(root.Val) > val {
		if float64(root.Val)-val < *ans {
			*ans = float64(root.Val) - val
			*v = root.Val
		}
	} else {
		if val-float64(root.Val) < *ans {
			*ans = val - float64(root.Val)
			*v = root.Val
		}
	}
	_closest_val(root.Right, val, ans, v)
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(5)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	fmt.Println(closest_val(root, 3.714286))
}
