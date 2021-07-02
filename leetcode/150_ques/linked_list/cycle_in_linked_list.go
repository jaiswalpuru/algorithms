/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
*/

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func New(val int) *Node { return &Node{data: val, next: nil} }

func is_cycle_present(root *Node) bool {
	sp, fp := root, root

	for sp != nil && fp != nil && fp.next != nil {
		sp = sp.next
		fp = fp.next.next
		if sp == fp {
			return true
		}
	}
	return false
}

func main() {
	a := New(3)
	a.next = New(2)
	a.next.next = New(0)
	a.next.next.next = New(-4)
	a.next.next.next.next = a.next
	fmt.Println(is_cycle_present(a))
	b := New(1)
	b.next = New(2)
	b.next.next = b
	fmt.Println(is_cycle_present(b))
	c := New(1)
	fmt.Println(is_cycle_present(c))
}
