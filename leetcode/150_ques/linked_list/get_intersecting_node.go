/*
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.
For example, the following two linked lists begin to intersect at node c1:
It is guaranteed that there are no cycles anywhere in the entire linked structure.
Note that the linked lists must retain their original structure after the function returns.

Example 1:

Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'
Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,6,1,8,4,5].
There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.


Example 2:

Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'
Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [1,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes
before the intersected node in A; There are 1 node before the intersected node in B.


Example 3:

Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection
Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5].
Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.
*/

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func New(data int) *Node { return &Node{data: data, next: nil} }

func length(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + length(root.next)
}

func intersecting_node(a, b *Node) *Node {
	l1, l2 := a, b
	m, n := length(l1), length(l2)
	if m > n {
		for i := 0; i < (m - n); i++ {
			l1 = l1.next
		}
	} else {
		for i := 0; i < (n - m); i++ {
			l2 = l2.next
		}
	}

	for l1 != l2 {
		l1 = l1.next
		l2 = l2.next
	}
	return l1
}

func main() {
	list_a := New(4)
	list_a.next = New(1)
	list_a.next.next = New(8)
	list_a.next.next.next = New(4)
	list_a.next.next.next.next = New(5)
	list_b := New(5)
	list_b.next = New(6)
	list_b.next.next = New(1)
	list_b.next.next.next = list_a.next.next
	fmt.Println(intersecting_node(list_a, list_b).data)
}
