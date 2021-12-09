/*
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
Input: p = [1,2], q = [1,null,2]
Output: false

Example 3:
Input: p = [1,2,1], q = [1,1,2]
Output: false
*/

package main

import "fmt"

type Node struct {
	Val         int
	Left, Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func is_same(root_one *Node, root_two *Node) bool {
	if root_one == nil && root_two == nil {
		return true
	} else if (root_one != nil && root_two == nil) || (root_one == nil && root_two != nil) {
		return false
	} else if root_one.Val != root_two.Val {
		return false
	} else {
		return is_same(root_one.Left, root_two.Left) && is_same(root_one.Right, root_two.Right)
	}
}

func main() {
	r1 := New(1)
	r1.Left = New(2)
	r1.Right = New(3)
	r2 := New(1)
	r2.Left = New(2)
	r2.Right = New(3)
	fmt.Println(is_same(r1, r2))
}
