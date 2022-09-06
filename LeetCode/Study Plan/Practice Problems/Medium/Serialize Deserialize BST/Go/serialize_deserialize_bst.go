package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Codec struct{}

func pre_order(root *TreeNode, res *string) {
	if root == nil {
		return
	}
	*res += strconv.Itoa(root.Val) + " "
	pre_order(root.Left, res)
	pre_order(root.Right, res)
}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := ""
	pre_order(root, &res)
	return res[:len(res)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	node := strings.Split(data, " ")
	nodes := make([]int, len(node))
	for i := 0; i < len(node); i++ {
		nodes[i], _ = strconv.Atoi(node[i])
	}
	root := &TreeNode{}
	for i, val := range nodes {
		if i == 0 {
			root.Val = val
		} else {
			insert(root, val)
		}
	}
	return root
}

func insert(root *TreeNode, val int) {
	if val <= root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val, Left: nil, Right: nil}
		} else {
			insert(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val, Left: nil, Right: nil}
		} else {
			insert(root.Right, val)
		}
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
