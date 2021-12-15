/*
Given a singly linked list and an integer k, remove the kth last element from the list.
k is guaranteed to be smaller than the length of the list.

The list is very long, so making more than one pass is prohibitively expensive.

Do this in constant space and in one pass.
*/

package main

import (
	"fmt"
	"strconv"
)

type LL struct {
	data int
	next *LL
}

func NewList(data int) *LL { return &LL{data: data, next: nil} }

func (l *LL) remove_kth_element_from_last(k int) {
	sp, fp := l, l
	for ; k > 0; k-- {
		fp = fp.next
	}
	for ; fp.next != nil; sp, fp = sp.next, fp.next {
	}
	sp.next = sp.next.next
}

func (l *LL) String() string {
	str := ""
	for l.next != nil {
		str += strconv.Itoa(l.data) + " -> "
		l = l.next
	}
	str += strconv.Itoa(l.data)
	return str
}

func main() {
	l := NewList(12)
	l.next = NewList(13)
	l.next.next = NewList(14)
	l.next.next.next = NewList(15)
	l.next.next.next.next = NewList(16)
	l.next.next.next.next.next = NewList(17)
	l.next.next.next.next.next.next = NewList(18)
	fmt.Println(l)
	k := 3 //remove the 3rd element from the last of the list in this case 15 should be deleted
	l.remove_kth_element_from_last(k)
	fmt.Println(l)
}
