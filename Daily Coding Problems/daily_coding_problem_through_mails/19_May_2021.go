/*
	Given the root to a binary search tree, find the second largest node in the tree.
	Here I used inorder to keep track of first largest and second largest.
	Another approach can be to insert the data into a max heap, and pop it two times
*/

package main

import (
	"fmt"
	"math"
)

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data} }

func inorder(root *Node, f, s *int) {
	if root == nil {
		return
	}
	inorder(root.left, f, s)
	if *f != int(math.MaxInt32) {
		if *f < root.data {
			*s = *f
			*f = root.data
		} else {
			if *s < root.data {
				*s = root.data
			}
		}
	}
	inorder(root.right, f, s)
}

func main() {
	root := New(15)
	root.left = New(9)
	root.right = New(20)
	root.left.left = New(5)
	root.left.right = New(12)
	root.right.right = New(25)
	root.right.left = New(17)
	root.right.right.right = New(30)
	root.right.right.right.right = New(40)
	first_largest, sec_largest := int(math.MinInt32), int(math.MinInt32)
	inorder(root, &first_largest, &sec_largest)
	fmt.Println(first_largest, sec_largest)
}
