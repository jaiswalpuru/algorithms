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
}
