/*
You are given the root node of a binary search tree (BST) and a value to insert into the tree.
Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion.
You can return any of them.
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var new_line = fmt.Println

func New(Val int) *TreeNode { return &TreeNode{Val: Val, Left: nil, Right: nil} }

func insert_non_recursive(root *TreeNode, val int) *TreeNode {
	temp := root
	var prev *TreeNode
	for {
		prev = temp
		if temp.Val > val {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
		if temp == nil {
			break
		}
	}
	if prev.Val > val {
		prev.Left = New(val)
	} else {
		prev.Right = New(val)
	}
	return root
}

//-------------------------Recursive---------------------------
func insert_recursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return New(val)
	}
	if root.Val > val {
		root.Right = insert_recursive(root.Right, val)
	} else {
		root.Left = insert_non_recursive(root.Left, val)
	}
	return root
}

//-------------------------Recursive---------------------------

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%d->", root.Val)
	inorder(root.Right)
}

func main() {
	root := New(4)
	root.Left = New(2)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	root = insert_non_recursive(root, 5)
	inorder(root)
	new_line()
}
