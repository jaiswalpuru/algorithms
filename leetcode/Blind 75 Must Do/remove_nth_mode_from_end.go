/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]
*/

package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data int
	next *Node
}

func New(val int) *Node { return &Node{data: val, next: nil} }

func (r *Node) String() string {
	str := ""
	temp := r
	for temp != nil {
		str += strconv.Itoa(temp.data) + "->"
		temp = temp.next
	}
	return str
}

func remove_nth_node_from_end(root *Node, n int) *Node {
	sp, fp := root, root

	if root == nil {
		return root
	}

	for n != 0 {
		fp = fp.next
		n--
	}
	if fp == nil {
		return fp
	}
	var prev *Node
	for fp != nil {
		prev = sp
		sp = sp.next
		fp = fp.next
	}
	prev.next = sp.next
	return root
}

func main() {
	root := New(1)
	root.next = New(2)
	root.next.next = New(3)
	root.next.next.next = New(4)
	root.next.next.next.next = New(5)
	fmt.Println(remove_nth_node_from_end(root, 2))
	root_new := New(1)
	fmt.Println(remove_nth_node_from_end(root_new, 1))
	root_new = New(1)
	root_new.next = New(2)
	fmt.Println(remove_nth_node_from_end(root_new, 1))
}
