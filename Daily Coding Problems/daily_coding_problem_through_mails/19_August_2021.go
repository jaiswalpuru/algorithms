/*
Let's represent an integer in a linked list format by having each node represent a digit in the number.
The nodes make up the number in reversed order.

For example, the following linked list:

1 -> 2 -> 3 -> 4 -> 5
is the number 54321.

Given two linked lists in this format, return their sum in the same linked list format.

For example, given

9 -> 9
5 -> 2
return 124 (99 + 25) as:

4 -> 2 -> 1
*/

package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func New(val int) *Node { return &Node{val: val, next: nil} }

//adding the list of same size
func add_list_same_size(r1, r2 *Node, sum, carry *int, res *[]int) {

	if r1 == nil && r2 == nil {
		return
	}

	if r1 == nil && r2 != nil {
		*sum += r2.val
	} else if r1 != nil && r2 == nil {
		*sum += r1.val
	} else {
		*sum += r1.val + r2.val
	}
	*sum += *carry
	*carry = *sum / 10
	*sum = *sum % 10
	*res = append(*res, *sum)
	*sum = 0
	if r1 != nil && r2 != nil {
		add_list_same_size(r1.next, r2.next, sum, carry, res)
	} else if r1 != nil && r2 == nil {
		add_list_same_size(r1.next, nil, sum, carry, res)
	} else if r1 == nil && r2 != nil {
		add_list_same_size(nil, r2.next, sum, carry, res)
	}
}

func add(r1, r2 *Node) []int {
	sum, carry := 0, 0
	res := []int{}
	add_list_same_size(r1, r2, &sum, &carry, &res)
	if carry != 0 {
		res = append(res, carry)
	}
	return res
}

//The result which is printed here is in the reverse order
func main() {
	r1 := New(3)
	r1.next = New(6)
	r1.next.next = New(5)
	r2 := New(2)
	r2.next = New(4)
	fmt.Println(add(r1, r2))
}
