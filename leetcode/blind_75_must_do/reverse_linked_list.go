/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []
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

func reverse_ll(root *Node) *Node {
	if root == nil {
		return nil
	}
	temp := reverse_ll(root.next)
	if temp == nil {
		return root
	}
	root.next.next = root
	return temp
}

func (root *Node) String() string {
	head := root
	str := ""
	for head != nil {
		if head.next == nil {
			str += strconv.Itoa(head.data)
		} else {
			str += strconv.Itoa(head.data) + " -> "
		}
		head = head.next
	}
	return fmt.Sprintf("%v", str)
}

func main() {
	root := New(1)
	root.next = New(2)
	root.next.next = New(3)
	root.next.next.next = New(4)
	fmt.Println(root)
	head := reverse_ll(root)
	root.next = nil
	fmt.Println(head)
}
