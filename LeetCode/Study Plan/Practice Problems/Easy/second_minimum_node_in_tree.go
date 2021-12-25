/*
Given a non-empty special binary tree consisting of nodes with the non-negative value,
where each node in this tree has exactly two or zero sub-node. If the node has two sub-nodes,
then this node's value is the smaller value among its two sub-nodes. More formally, the property
root.val = min(root.left.val, root.right.val) always holds.

Given such a binary tree, you need to output the second minimum value in the set made of all the nodes'
value in the whole tree.

If no such second minimum value exists, output -1 instead.
*/

package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func New(val int) *Node { return &Node{Val: val, Left: nil, Right: nil} }

func inorder(node *Node, u *int, v *int) {
	if node == nil {
		return
	}
	inorder(node.Left, u, v)
	if node.Val > *u && node.Val < *v {
		*v = node.Val
	}
	inorder(node.Right, u, v)
}

func sec_min(node *Node) int {
	u, v := node.Val, int(math.MaxInt64)
	inorder(node, &u, &v)
	if v == int(math.MaxInt64) {
		return -1
	}
	return v
}

func main() {
	node := New(2)
	node.Left = New(2)
	node.Right = New(5)
	node.Right.Left = New(5)
	node.Right.Right = New(7)
	fmt.Println(sec_min(node))
}
