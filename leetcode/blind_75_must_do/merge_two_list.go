/*
Merge two sorted linked lists and return it as a sorted list.
The list should be made by splicing together the nodes of the first two lists.

Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: l1 = [], l2 = []
Output: []

Example 3:
Input: l1 = [], l2 = [0]
Output: [0]
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

func New(data int) *Node { return &Node{data: data, next: nil} }

func move_node(dst **Node, src **Node) {
	new_node := *src
	*src = new_node.next
	new_node.next = *dst
	*dst = new_node
}

func merge_list(l1, l2 *Node) *Node {
	var res *Node = nil
	var last_ref **Node = &res

	if l1 == nil {
		*last_ref = l2
	}
	if l2 == nil {
		*last_ref = l1
	}

	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			move_node(last_ref, &l1)
		} else {
			move_node(last_ref, &l2)
		}
		last_ref = &((*last_ref).next)
	}
	if l1 != nil {
		*last_ref = l1
	}
	if l2 != nil {
		*last_ref = l2
	}

	return res
}

func (l *Node) String() string {
	str := ""
	for l != nil {
		str += strconv.Itoa(l.data) + "->"
		l = l.next
	}
	return str
}

func main() {
	l1 := New(1)
	l1.next = New(2)
	l1.next.next = New(4)
	l2 := New(1)
	l2.next = New(3)
	l2.next.next = New(4)
	fmt.Println(merge_list(l1, l2))
}
