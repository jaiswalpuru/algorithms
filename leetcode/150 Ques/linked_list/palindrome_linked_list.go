/*
Given the head of a singly linked list, return true if it is a palindrome.

Example 1:


Input: head = [1,2,2,1]
Output: true


Example 2:

Input: head = [1,2]
Output: false
*/

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func New(data int) *Node { return &Node{data: data, next: nil} }

func is_palindrome_stack(root *Node) bool {
	stack := []int{}
	sp := root
	for sp != nil {
		stack = append(stack, sp.data)
		sp = sp.next
	}

	sp = root
	for sp != nil {
		n := len(stack)
		if stack[n-1] != sp.data {
			return false
		}
		stack = stack[:n-1]
		sp = sp.next
	}

	if len(stack) == 0 && sp != nil {
		return false
	}

	return true
}

func is_palindrome_recursive(left **Node, right *Node) bool {

	if right == nil {
		return true
	}

	is_palin := is_palindrome_recursive(left, right.next)
	if !is_palin {
		return false
	}

	is_palin_one := right.data == (*left).data
	*left = (*left).next
	return is_palin_one
}

func is_palindrome(root *Node) bool { return is_palindrome_recursive(&root, root) }

func main() {
	l1 := New(1)
	l1.next = New(2)
	l1.next.next = New(2)
	l1.next.next.next = New(1)
	fmt.Println(is_palindrome(l1))
	l2 := New(1)
	l2.next = New(2)
	fmt.Println(is_palindrome_stack(l2))
}
