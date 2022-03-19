package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val : val, Next : nil} }

func is_palindrome(head *ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	l,r := 0,len(arr)-1
	for l<=r{
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func is_palindrome_recursive(tail *ListNode, temp **ListNode) bool {
	if tail != nil {
		if !is_palindrome_recursive(tail.Next, temp) { 
			return false 
		}

		if tail.Val != (*temp).Val {
			return false
		}
		*temp = (*temp).Next
	}
	return true
}

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(2)
	head.Next.Next.Next = New(1)
	// fmt.Println(is_palindrome(head))
	temp := head
	fmt.Println(is_palindrome_recursive(head, &temp))
}