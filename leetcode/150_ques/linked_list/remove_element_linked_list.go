/*
Given the head of a linked list and an integer val, remove all the nodes of the linked
list that has Node.val == val, and return the new head.

Example 1:

Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]


Example 2:

Input: head = [], val = 1
Output: []


Example 3:

Input: head = [7,7,7,7], val = 7
Output: []
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

func (r *Node) String() string {
	str := ""
	temp := r
	for temp != nil {
		str += strconv.Itoa(temp.val) + "->"
		temp = temp.next
	}
	return str
}

func remove_elements(root *Node, val int) {
	prev, cur := root, root
	if cur == nil {
		return
	}

	for cur != nil {
		if cur.val == val {
			prev.next = cur.next
		}
		prev = cur
		cur = cur.next
	}
}

func main() {
	p := New(1)
	p.next = New(2)
	p.next.next = New(6)
	p.next.next.next = New(3)
	p.next.next.next.next = New(4)
	p.next.next.next.next.next = New(5)
	p.next.next.next.next.next.next = New(6)
	remove_elements(p, 6)
	fmt.Println(p)
}
