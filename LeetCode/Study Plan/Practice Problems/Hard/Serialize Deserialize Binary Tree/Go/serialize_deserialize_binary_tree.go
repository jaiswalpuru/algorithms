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

func Constructor() Codec { return Codec{} }

func preorder(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
		return str
	}
	str += strconv.Itoa(root.Val) + ","
	str = preorder(root.Left, str)
	str = preorder(root.Right, str)
	return str
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string { return preorder(root, "") }

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	return reserialize(&arr)
}

func reserialize(pre_order *[]string) *TreeNode {
	if (*pre_order)[0] == "null" {
		*pre_order = (*pre_order)[1:]
		return nil
	}
	val, _ := strconv.Atoi((*pre_order)[0])
	*pre_order = (*pre_order)[1:]
	root := New(val)
	root.Left = reserialize(pre_order)
	root.Right = reserialize(pre_order)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
