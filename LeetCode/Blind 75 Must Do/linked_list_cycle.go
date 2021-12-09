/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by
continuously following the next pointer. Internally, pos is used to denote the index of the node
that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
*/

package main

import "fmt"

type LL struct {
	val  int
	next *LL
}

func New(data int) *LL { return &LL{val: data, next: nil} }

func has_cycle(head *LL) bool {
	if head == nil {
		return false
	}
	sp, fp := head, head
	for fp != nil && fp.next != nil {
		sp = sp.next
		fp = fp.next.next
		if sp == fp {
			return true
		}
	}
	return false
}

func main() {
	head := New(3)
	head.next = New(2)
	head.next.next = New(0)
	head.next.next.next = New(-4)
	head.next.next.next.next = head.next
	fmt.Println(has_cycle(head))
}
