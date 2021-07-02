/*
Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem
without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Example 1:

Input: head = [1,2,3,4]
Output: [2,1,4,3]


Example 2:

Input: head = []
Output: []
Example 3:

Input: head = [1]
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

func New(data int) *Node { return &Node{data: data, next: nil} }

func swap_pair(s *Node) *Node {
	if s == nil || s.next == nil {
		return s
	}
	remaining := s.next.next
	new_head := s.next
	s.next.next = s
	s.next = swap_pair(remaining)
	return new_head
}

func (s *Node) String() string {
	str := ""
	temp := s
	for temp != nil {
		str += strconv.Itoa(temp.data) + "->"
		temp = temp.next
	}
	return str
}

func main() {
	s := New(1)
	s.next = New(2)
	s.next.next = New(3)
	s.next.next.next = New(4)
	s = swap_pair(s)
	fmt.Println(s)
}
