package bst

import "fmt"

type BinarySearchTreeNode struct {
	data int
	left,
	right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func NewNode() *BinarySearchTree { return &BinarySearchTree{} }

func Inorder(root *BinarySearchTreeNode) {
	if root != nil {
		Inorder(root.left)
		fmt.Print(root.data, " ")
		Inorder(root.right)
	}
}

func Preorder(root *BinarySearchTreeNode) {
	if root != nil {
		fmt.Println(root.data, " ")
		Preorder(root.left)
		Preorder(root.right)
	}
}

func Postorder(root *BinarySearchTreeNode) {
	if root != nil {
		Postorder(root.left)
		Postorder(root.right)
		fmt.Println(root.data, " ")
	}
}

func (root *BinarySearchTree) TreeMinimum() *BinarySearchTreeNode {
	x := root.root
	for x != nil {
		x = x.left
	}
	return x
}

func (root *BinarySearchTree) TreeMaximum() *BinarySearchTreeNode {
	x := root.root
	for x != nil {
		x = x.right
	}
	return x
}

func TreeSearch(root *BinarySearchTreeNode, val int) *BinarySearchTreeNode {
	if root == nil || root.data == val {
		return root
	}
	if val > root.data {
		return TreeSearch(root.right, val)
	}
	return TreeSearch(root.right, val)
}

func TreeSearchIterative(root *BinarySearchTreeNode, val int) *BinarySearchTreeNode {
	for root != nil {
		if root.data == val {
			return root
		}
		if val > root.data {
			root = root.right
		} else {
			root = root.left
		}
	}
	return root
}

func (T *BinarySearchTree) Insert(val int) {

}
