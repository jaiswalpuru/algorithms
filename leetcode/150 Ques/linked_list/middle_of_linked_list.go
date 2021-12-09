/*
Given a non-empty, singly linked list with head node head, return a middle node of linked list.

If there are two middle nodes, return the second middle node.



Example 1:

Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
Example 2:

Input: [1,2,3,4,5,6]
Output: Node 4 from this list (Serialization: [4,5,6])
Since the list has two middle nodes with values 3 and 4, we return the second one.
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

func middle_element(root *Node) int {
	if root == nil {
		return -1
	}
	sp, fp := root, root
	for fp != nil && fp.next != nil {
		fmt.Println(sp.val, fp.val)
		sp = sp.next
		fp = fp.next.next
	}
	return sp.val
}

func main() {
	p := New(1)
	p.next = New(2)
	p.next.next = New(3)
	p.next.next.next = New(4)
	p.next.next.next.next = New(5)
	p.next.next.next.next.next = New(6)
	fmt.Println(middle_element(p))
}
