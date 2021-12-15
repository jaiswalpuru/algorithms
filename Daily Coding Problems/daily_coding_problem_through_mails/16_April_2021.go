//Given the root to a binary tree, implement serialize(root),
//which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MARKER = "-1"
)

type BinaryTree struct {
	data        int
	left, right *BinaryTree
}

func NewBT(data int) *BinaryTree {
	bt := &BinaryTree{data: data, left: nil, right: nil}
	return bt
}

func insertNode(root *BinaryTree, data int) *BinaryTree {
	if root == nil {
		return NewBT(data)
	}
	if data < root.data {
		root.left = insertNode(root.left, data)
		return root
	}
	root.right = insertNode(root.right, data)
	return root
}

func serialize(root *BinaryTree, ans *string) {
	if root == nil {
		*ans += MARKER + " "
		return
	}
	*ans += strconv.Itoa(root.data) + " "
	serialize(root.left, ans)
	serialize(root.right, ans)
}

func deserialize(serializedTree []int) *BinaryTree {
	bt := NewBT(serializedTree[0])
	for i := 1; i < len(serializedTree); i++ {
		insertNode(bt, serializedTree[i])
	}
	return bt
}

func main() {
	bt := NewBT(20)
	bt.left = NewBT(8)
	bt.right = NewBT(22)
	bt.left.left = NewBT(4)
	bt.left.right = NewBT(12)
	bt.left.right.left = NewBT(10)
	bt.left.right.right = NewBT(14)
	serializeTree := ""
	serialize(bt, &serializeTree)
	fmt.Println(serializeTree)
	serializedTree := strings.Split(serializeTree, " ")
	treeVals := []int{}
	for i := 0; i < len(serializedTree); i++ {
		if serializedTree[i] != "-1" {
			temp, _ := strconv.Atoi(serializedTree[i])
			treeVals = append(treeVals, temp)
		}
	}
	root1 := deserialize(treeVals)
	fmt.Println(root1.left.left)
}
