/*
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
If the two linked lists have no intersection at all, return null.

Note that the linked lists must retain their original structure after the function returns.

Custom Judge:

The inputs to the judge are given as follows (your program is not given these inputs):

intersectVal - The value of the node where the intersection occurs. This is 0 if there is no intersected node.
listA - The first linked list.
listB - The second linked list.
skipA - The number of nodes to skip ahead in listA (starting from the head) to get to the intersected node.
skipB - The number of nodes to skip ahead in listB (starting from the head) to get to the intersected node.
The judge will then create the linked structure based on these inputs and pass the two heads, headA and headB to your program.
If you correctly return the intersected node, then your solution will be accepted.
*/

/*
One more approach to this problem would be to use two pointers, calculate the length of both lists and place pointer one to the
longer list and pointer b to the shorter list them move one step at a time checking equality
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func get_intersection_node(a, b *ListNode) *ListNode {
	mp := make(map[*ListNode]struct{})
	for a != nil {
		mp[a] = struct{}{}
		a = a.Next
	}

	for b != nil {
		if _, ok := mp[b]; ok {
			return b
		}
		b = b.Next
	}

	return nil
}

func main() {
	a := New(4)
	a.Next = New(1)
	a.Next.Next = New(8)
	a.Next.Next.Next = New(4)
	a.Next.Next.Next.Next = New(5)
	b := New(5)
	b.Next = New(6)
	b.Next.Next = New(1)
	b.Next.Next.Next = a.Next.Next
	fmt.Println(get_intersection_node(a, b))
}
