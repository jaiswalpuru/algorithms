package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func pseudo_palindromic_path_in_binary_tree(root *TreeNode) int {
	cnt := 0
	preorder(root, map[int]int{}, &cnt)
	return cnt
}

func preorder(root *TreeNode, mp map[int]int, cnt *int) {
	if root == nil {
		return
	}
	mp[root.Val]++
	if root.Left == nil && root.Right == nil {
		if check_palindrome(mp) {
			*cnt++
		}
	}
	preorder(root.Left, mp, cnt)
	preorder(root.Right, mp, cnt)
	mp[root.Val]--
}

func check_palindrome(hash_map map[int]int) bool {
	cnt := 0
	for _, v := range hash_map {
		if v%2 != 0 {
			cnt++
		}
	}
	return cnt <= 1
}

func main() {
	root := New(2)
	root.Left = New(3)
	root.Right = New(1)
	root.Left.Left = New(3)
	root.Left.Right = New(1)
	root.Right.Right = New(1)
	fmt.Println(pseudo_palindromic_path_in_binary_tree(root))
}
