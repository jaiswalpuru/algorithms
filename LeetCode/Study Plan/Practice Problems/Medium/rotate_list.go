/*
Given the head of a linked list, rotate the list to the right by k places.
*/

package main

import "fmt"

type Node struct {
	val  int
	Next *Node
}

func New(val int) *Node { return &Node{val: val, Next: nil} }

func get_list_length(head *Node) int {
	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}
	return cnt
}

func rotate_list(head *Node, k int) *Node {
	n := get_list_length(head)
	if n == 0 || n == 1 || k == 0 {
		return head
	}
	k = k % n
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = head

	temp = head
	k = n - k
	var prev *Node
	for k > 0 {
		prev = temp
		temp = temp.Next
		k--
	}
	prev.Next = nil
	head = temp
	return head
}

func main() {
	head := New(1)
	head.Next = New(2)
	// head.next.next = New(3)
	// head.next.next.next = New(4)
	// head.next.next.next.next = New(5)
	fmt.Println(rotate_list(head, 1))
}
