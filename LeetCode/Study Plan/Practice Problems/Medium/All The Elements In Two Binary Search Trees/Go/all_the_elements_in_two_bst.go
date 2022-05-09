package main

import "fmt"
import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func inorder(r *TreeNode, res *[]int) {
	if r == nil {
		return
	}
	inorder(r.Left, res)
	*res = append(*res, r.Val)
	inorder(r.Right, res)
}

func all_elements_in_two_bst(r1, r2 *TreeNode) []int {
	r_1 := []int{}
	r_2 := []int{}
	res := []int{}
	inorder(r1, &r_1)
	inorder(r2, &r_2)
	res = append(res, r_1...)
	res = append(res, r_2...)
	sort.Ints(res)
	return res
}

func main() {
	root := New(2)
	root.Left = New(1)
	root.Right = New(4)
	root_t := New(1)
	root_t.Left = New(0)
	root_t.Right = New(3)
	fmt.Println(all_elements_in_two_bst(root, root_t))
}
