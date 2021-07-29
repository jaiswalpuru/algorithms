/*
Print the nodes in a binary tree level-wise. For example, the following should print 1, 2, 3, 4, 5.

  1
 / \
2   3
   / \
  4   5
*/

package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data, left: nil, right: nil} }

func print_level_order(root *Node) {
	if root == nil {
		return
	}

	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		fmt.Print(cur.data, "->")
		if cur.left != nil {
			q = append(q, cur.left)
		}
		if cur.right != nil {
			q = append(q, cur.right)
		}
	}
	fmt.Println()
}

func main() {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.right.left = New(4)
	root.right.right = New(5)
	print_level_order(root)
}
