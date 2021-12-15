/*
Given the root of a binary tree, return a deepest node. For example, in the following tree, return d.

    a
   / \
  b   c
 /
d
*/

package main

import "fmt"

type Node struct {
	data  rune
	left  *Node
	right *Node
}

func New(data rune) *Node { return &Node{data: data, left: nil, right: nil} }

func deepest_node(root *Node) rune {
	if root == nil {
		return []rune("")[0]
	}
	queue := []*Node{root}
	var node *Node
	for len(queue) > 0 {
		q_len := len(queue)
		var lvl []rune
		for i := 0; i < q_len; i++ {
			node = queue[0]
			lvl = append(lvl, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return node.data
}

func main() {
	root := New('a')
	root.left = New('b')
	root.right = New('c')
	root.left.left = New('d')
	fmt.Println(string(deepest_node(root)))
}
