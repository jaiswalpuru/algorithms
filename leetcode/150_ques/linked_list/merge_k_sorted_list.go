/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
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
Output: []
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

func New(val int) *Node { return &Node{val: val, next: nil} }

func (r *Node) String() string {
	str := ""
	temp := r
	for temp != nil {
		str += strconv.Itoa(temp.val) + "->"
		temp = temp.next
	}
	return str
}

func merge_list_recursive(l1, l2 *Node) *Node {

	result := new(Node)

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.val <= l2.val {
		result = l1
		result.next = merge_list_recursive(l1.next, l2)
	} else {
		result = l2
		result.next = merge_list_recursive(l1, l2.next)
	}
	return result
}

func merge_sorted_list(list []*Node, r int) *Node {
	for r != 0 {
		i, j := 0, r
		for i < j {
			list[i] = merge_list_recursive(list[i], list[j])
			i++
			j--
			if i >= j {
				r = j
			}
		}
	}
	return list[0]
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
	fmt.Println(merge_sorted_list(list, 2))
}
