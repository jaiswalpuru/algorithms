package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

//-------------------Brute Force (Accepted)-----------------------
func next_greater_ele(head *ListNode) []int {
	if head == nil {
		return nil
	}

	ans := []int{}
	for head != nil {
		ans = append(ans, _next_greater_ele(head.Next, head.Val))
		head = head.Next
	}
	return ans
}

func _next_greater_ele(head *ListNode, val int) int {
	if head == nil {
		return 0
	}
	for head != nil && val >= head.Val {
		head = head.Next
	}
	if head == nil {
		return 0
	}
	return head.Val
}

//----------------------------------------------------------------

type Pair struct{ val, idx int }

func get_list_len(head *ListNode) int {
	n := 0
	for head != nil {
		head = head.Next
		n++
	}
	return n
}

//------------------using monotonic stack-------------------------
func next_greater_ele_eff(head *ListNode) []int {
	if head == nil {
		return nil
	}

	ans := make([]int, get_list_len(head))
	i := 0
	stack := []Pair{{val: head.Val, idx: i}}
	for head != nil {
		for len(stack) > 0 && stack[len(stack)-1].val < head.Val {
			_, idx := stack[len(stack)-1].val, stack[len(stack)-1].idx
			stack = stack[:len(stack)-1]
			ans[idx] = head.Val
		}
		stack = append(stack, Pair{val: head.Val, idx: i})
		head = head.Next
		i++
	}
	return ans
}

//------------------using monotonic stack-------------------------

func main() {
	head := New(2)
	head.Next = New(7)
	head.Next.Next = New(4)
	head.Next.Next.Next = New(3)
	head.Next.Next.Next.Next = New(5)
	fmt.Println(next_greater_ele_eff(head))
}
