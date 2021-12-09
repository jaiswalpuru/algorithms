/*You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.



Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
Example 2:

Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []*/

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

func merge_list_recursive(l1, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *Node
	if l1.data <= l2.data {
		result = l1
		result.next = merge_list_recursive(l1.next, l2)
	} else {
		result = l2
		result.next = merge_list_recursive(l1, l2.next)
	}
	return result
}

func (l *Node) String() string {
	str := ""
	for l != nil {
		str += strconv.Itoa(l.data) + "->"
		l = l.next
	}
	return str
}

func merge_k_sorted(l []*Node, r int) *Node {
	for r != 0 {
		i, j := 0, r
		for i < j {
			l[i] = merge_list_recursive(l[i], l[j])
			i++
			j--
			if i >= j {
				r = j
			}
		}
	}
	return l[0]
}

func main() {
	list := make([]*Node, 3)
	list[0] = New(1)
	list[0].next = New(4)
	list[0].next.next = New(5)
	list[1] = New(1)
	list[1].next = New(3)
	list[1].next.next = New(4)
	list[2] = New(2)
	list[2].next = New(6)
	for i := 0; i < 3; i++ {
		fmt.Println(list[i])
	}
	fmt.Println(merge_k_sorted(list, 2))
}
