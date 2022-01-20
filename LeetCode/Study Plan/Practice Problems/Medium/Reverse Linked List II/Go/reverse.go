package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var left *ListNode
var stop bool

func reverse(node *ListNode, m, n int) *ListNode {
	stop = false
	left = node
	recurse_reverse(node, m, n)
	return node
}

//refer lc soln
func recurse_reverse(right *ListNode, m, n int) {
	if n == 1 {
		return
	}
	right = right.Next

	//the point from where recursion will start
	if m > 1 {
		left = left.Next
	}

	recurse_reverse(right, m-1, n-1)

	//if both pointers become equal or cross each other then stop
	if left == right || right.Next == left {
		stop = true
	}

	if !stop {
		val := left.Val
		left.Val = right.Val
		right.Val = val

		//the right will move via backtracking
		left = left.Next
	}

}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func (l *ListNode) String() string {
	str := ""
	for l != nil {
		if l.Next == nil {
			str += strconv.FormatInt(int64(l.Val), 10)
		} else {
			str += strconv.FormatInt(int64(l.Val), 10) + "->"
		}
		l = l.Next
	}
	return str
}

func iterative_reverse(node *ListNode, m, n int) *ListNode {
	if node == nil {
		return node
	}

	var prev *ListNode = nil
	curr := node
	for m > 1 {
		prev = curr
		curr = curr.Next
		m--
		n--
	}

	con, tail := prev, curr

	var temp *ListNode
	for n > 0 {
		temp = curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
		n--
	}
	if con != nil {
		con.Next = prev
	} else {
		node = prev
	}
	tail.Next = curr
	return node
}

func main() {
	node := New(1)
	node.Next = New(2)
	node.Next.Next = New(3)
	// node.Next.Next.Next = New(4)
	// node.Next.Next.Next.Next = New(5)
	fmt.Println(iterative_reverse(node, 1, 2))
}
