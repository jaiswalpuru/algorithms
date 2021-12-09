/*
Serialization is converting a data structure or object into a sequence of bits so that it can
be stored in a file or memory buffer, or transmitted across a network connection link to be
reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary search tree. There is no restriction
on how your serialization/deserialization algorithm should work. You need to ensure that a binary
search tree can be serialized to a string, and this string can be deserialized to the original tree structure.

The encoded string should be as compact as possible.

Example 1:
Input: root = [2,1,3]
Output: [2,1,3]

Example 2:
Input: root = []
Output: []
*/

package main

import (
	"fmt"
	"strconv"
)

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec { return Codec{} }

var d []int

func insert(root *TreeNode, val int) {
	if val <= root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
		} else {
			insert(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
		} else {
			insert(root.Right, val)
		}
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var pre_order []int
	_pre_order(root, &pre_order)
	str, n := "", len(pre_order)
	for i := 0; i < n; i++ {
		str += strconv.Itoa(pre_order[i])
	}
	d = pre_order
	return str
}

func _pre_order(root *TreeNode, ele *[]int) {
	if root == nil {
		return
	}
	*ele = append(*ele, root.Val)
	_pre_order(root.Left, ele)
	_pre_order(root.Right, ele)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	root := &TreeNode{}
	for i, n := range d {
		if i == 0 {
			root.Val = n
		} else {
			insert(root, n)
		}
	}
	return root
}

func main() {
	root := &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	ser := Constructor()
	deser := Constructor()
	tree := ser.serialize(root)
	ans := deser.deserialize(tree)
	fmt.Println(ans)
}
