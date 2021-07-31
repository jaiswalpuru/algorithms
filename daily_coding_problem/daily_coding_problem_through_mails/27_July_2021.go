/*
Determine whether a doubly linked list is a palindrome. What if it’s singly linked?

For example, 1 -> 4 -> 3 -> 4 -> 1 returns True while 1 -> 4 returns False.
*/

package main

import "fmt"

type DLL struct {
	left, right *DLL
	data        int
}

func New(data int) *DLL { return &DLL{data: data, left: nil, right: nil} }

func check_palindrome(head *DLL) bool {
	p1, p2 := head, head
	for p2.right != nil {
		p2 = p2.right
	}
	for p1 != p2 {
		if p1.data != p2.data {
			return false
		}
		p1 = p1.right
		p2 = p2.left
	}
	return true
}

type LL struct {
	data int
	next *LL
}

func New_LL(data int) *LL { return &LL{data: data, next: nil} }

func check_palindrome_ll(l *LL) bool {

	sp, fp := l, l
	stack := []*LL{}
	for fp != nil && fp.next != nil {
		stack = append(stack, sp)
		sp = sp.next
		fp = fp.next.next
	}
	if fp != nil {
		sp = sp.next
	}
	for len(stack) > 0 && sp != nil {
		n := len(stack)
		val := stack[n-1]
		stack = stack[:n-1]
		if val.data != sp.data {
			return false
		}
		sp = sp.next
	}
	if (sp == nil && len(stack) > 0) || (sp != nil && len(stack) <= 0) {
		return false
	}
	return true
}

func main() {
	root := New(1)
	root.right = New(4)
	root.right.left = root
	root.right.right = New(3)
	root.right.right.left = root.right
	root.right.right.right = New(4)
	root.right.right.right.left = root.right.right
	root.right.right.right.right = New(1)
	root.right.right.right.right.left = root.right.right.right
	fmt.Println(check_palindrome(root))
	root_one := New(1)
	root_one.right = New(4)
	root_one.right.left = root
	fmt.Println(check_palindrome(root_one))
	root_ll := New_LL(1)
	root_ll.next = New_LL(4)
	root_ll.next.next = New_LL(3)
	root_ll.next.next.next = New_LL(4)
	root_ll.next.next.next.next = New_LL(1)
	fmt.Println(check_palindrome_ll(root_ll))
}
