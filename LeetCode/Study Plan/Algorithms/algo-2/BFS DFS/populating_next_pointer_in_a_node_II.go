/*
Given a binary tree

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node,
the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.
*/

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Pair struct {
	node *Node
	lvl  int
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := make([]Pair, 0)
	q = append(q, Pair{root, 0})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if len(q) == 0 {
			curr.node.Next = nil
		} else {
			if q[0].lvl == curr.lvl {
				curr.node.Next = q[0].node
			} else {
				curr.node.Next = nil
			}
		}

		if curr.node.Left != nil {
			q = append(q, Pair{curr.node.Left, curr.lvl + 1})
		}
		if curr.node.Right != nil {
			q = append(q, Pair{curr.node.Right, curr.lvl + 1})
		}
	}
	return root
}

func main() {
	root := &Node{Val: 1, Left: nil, Right: nil}
	root.Left = &Node{Val: 2, Left: nil, Right: nil}
	root.Right = &Node{Val: 3, Left: nil, Right: nil}
	root.Left.Left = &Node{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &Node{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &Node{Val: 7, Left: nil, Right: nil}
	fmt.Println(connect(root))
	fmt.Println(root.Left.Next)
}
