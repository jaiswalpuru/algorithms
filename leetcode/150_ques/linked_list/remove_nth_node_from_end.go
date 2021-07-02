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
	val  int
	next *Node
}

func New(data int) *Node { return &Node{val: data, next: nil} }

func delete_nth_node_from_end(root *Node, n int) *Node {
	sp, fp, prev := root, root, root
	//first move the fp to the nth node from start
	if fp == nil {
		return nil
	}
	for n > 0 {
		fp = fp.next
		n--
	}

	if fp == nil {
		return fp
	}

	for fp.next != nil {
		fp = fp.next
		prev = sp
		sp = sp.next
	}
	prev = sp
	sp = sp.next
	prev.next = sp.next

	return root
}

func (r *Node) String() string {
	str := ""
	temp := r
	for temp != nil {
		str += strconv.Itoa(temp.val) + "->"
		temp = temp.next
	}
	return str
}

func main() {
	root := New(1)
	root.next = New(2)
	root.next.next = New(3)
	root.next.next.next = New(4)
	root.next.next.next.next = New(5)
	fmt.Println(delete_nth_node_from_end(root, 2))
	node := New(1)
	fmt.Println(delete_nth_node_from_end(node, 1))
	node = New(1)
	node.next = New(2)
	fmt.Println(delete_nth_node_from_end(node, 1))
}
