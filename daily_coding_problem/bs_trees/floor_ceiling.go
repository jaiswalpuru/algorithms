/*
	Floor and ceiling value of a given integer.
	The floor is the highest element in the tree less than or equal to an integer, while the ceiling value is the
	lowest element in the tree greater than or equal to an integer.
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

func get_floor_ceil(root *Node, data, floor, ceil int) (int, int) {
	if root == nil {
		return floor, ceil
	}
	if root.data == data {
		return data, data
	}
	if root.data > data {
		floor, ceil = get_floor_ceil(root.left, data, floor, root.data)
	}
	if root.data < data {
		floor, ceil = get_floor_ceil(root.right, data, root.data, ceil)
	}
	return floor, ceil
}

func main() {
	//hardcoding binary search tree
	root := New(7)
	root.left = New(5)
	root.left.left = New(-1)
	root.left.right = New(6)
	root.right = New(10)
	root.right.right = New(25)
	fmt.Println(get_floor_ceil(root, 9, int(math.MinInt32), int(math.MinInt32)))
}
