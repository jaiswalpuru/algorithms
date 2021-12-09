/*
Write a function to delete a node in a singly-linked list.
You will not be given access to the head of the list, instead you will be given access to the node to be
deleted directly.

It is guaranteed that the node to be deleted is not a tail node in the list.
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

func delete_node(root *Node, val int) *Node {
	current, prev := root, root
	if current.val == val {
		prev = current.next
		return prev
	}
	for current != nil {
		if current.val == val {
			prev.next = current.next
		}
		prev = current
		current = current.next
	}
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
	node := New(4)
	node.next = New(5)
	node.next.next = New(1)
	node.next.next.next = New(9)
	fmt.Println(delete_node(node, 5))
}
