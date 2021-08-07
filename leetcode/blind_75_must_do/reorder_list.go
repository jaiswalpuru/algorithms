/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
*/

package main

import (
	"fmt"
	"strconv"
)

type LL struct {
	Val  int
	Next *LL
}

func New(val int) *LL { return &LL{Val: val, Next: nil} }

func (l *LL) String() string {
	str := ""
	head := l
	for head != nil {
		str += strconv.Itoa(head.Val) + "->"
		head = head.Next
	}
	return str
}

func reorder_list(l *LL) {

	if l == nil || l.Next == nil {
		return
	}

	sp, fp := l, l
	for fp != nil && fp.Next != nil {
		sp, fp = sp.Next, fp.Next.Next
	}
	fmt.Println(sp, fp)
	var prev *LL
	for sp != nil {
		sp.Next, prev, sp = prev, sp, sp.Next
	}

	f := l
	for prev.Next != nil {
		f.Next, f = prev, f.Next
		prev.Next, prev = f, prev.Next
	}
}

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(4)
	reorder_list(head)
	fmt.Println(head)
}
