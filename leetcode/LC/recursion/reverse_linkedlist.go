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

func New(data int) *Node { return &Node{data: data, next: nil} }

func reverse_list(root *Node) *Node {
	if root == nil {
		return nil
	}
	temp := reverse_list(root.next)
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
	node := New(1)
	node.next = New(2)
	node.next.next = New(3)
	node.next.next.next = New(4)
	node.next.next.next.next = New(5)
	fmt.Println(node)
	head := reverse_list(node)
	node.next = nil
	fmt.Println(head)
}
