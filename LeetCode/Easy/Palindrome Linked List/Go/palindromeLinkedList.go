package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	return checkPalindrome(head, &head)
}

func checkPalindrome(temp *ListNode, head **ListNode) bool {
	if temp != nil {
		if !checkPalindrome(temp.Next, head) {
			return false
		}
		if (*head).Val != temp.Val {
			return false
		}
		*head = (*head).Next
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
	fmt.Println(checkPalindrome(head, &temp))
}
