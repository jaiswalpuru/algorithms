/*
	Convert a sorted array to height-balance binary search tree
*/

package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data} }

func make_bst(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	root := New(arr[mid])
	root.left = make_bst(arr[:mid])
	root.right = make_bst(arr[mid+1:])
	return root
}

func inorder(root *Node, str *string) {
	if root == nil {
		return
	}
	inorder(root.left, str)
	*str += strconv.Itoa(root.data) + "->"
	inorder(root.right, str)
}

func (node *Node) String() string {
	str := ""
	inorder(node, &str)
	return fmt.Sprintf("%v", str)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	// the root will be the middle element if we want to create a balanced binary search tree
	fmt.Println(make_bst(arr))
}
