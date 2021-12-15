/*
	Given pre-order and in-order traversals reconstruct the tree
*/

package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data} }

func get_index(arr []int, val int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func reconstruct(pre, in []int) *Node {
	if pre == nil && in == nil {
		return nil
	}

	if len(pre) == 1 && len(in) == 1 {
		return &Node{data: pre[0]}
	}

	root := New(pre[0])
	root_ind := get_index(in, root.data)
	root.left = reconstruct(pre[1:1+root_ind], in[0:root_ind])
	root.right = reconstruct(pre[1+root_ind:], in[root_ind+1:])
	return root
}

func main() {
	in_order := []int{4, 2, 5, 1, 6, 3, 7}  //left->root->right
	pre_order := []int{1, 2, 4, 5, 3, 6, 7} //root->left->right

	fmt.Println("In-order : ", in_order)
	fmt.Println("Pre-order : ", pre_order)
	root := reconstruct(pre_order, in_order)
	validate(root)
	fmt.Println()
}

//prints the inorder traversal of the tree
func validate(root *Node) {
	if root == nil {
		return
	}
	validate(root.left)
	fmt.Print(root.data, " ")
	validate(root.right)
}
