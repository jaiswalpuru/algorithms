package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val} }

//--------------------------non recursive looks bad--------------------
func reverse_nodes_in_k_groups(head *ListNode, k int) *ListNode {
	st := []*ListNode{}
	root := New(-1)
	p := root
	for head != nil {
		k_temp := k
		for k_temp > 0 && head != nil {
			st = append(st, head)
			head = head.Next
			k_temp--
		}

		if len(st) == k {
			t := st[len(st)-1]
			st = st[:len(st)-1]
			temp := t
			for len(st) > 0 {
				temp.Next = st[len(st)-1]
				st = st[:len(st)-1]
				temp = temp.Next
			}
			temp.Next = nil
			if root.Val == -1 {
				root = t
				p = root
			} else {
				for p.Next != nil {
					p = p.Next
				}
				p.Next = t
			}
		} else {
			t := st[0]
			st = st[1:]
			temp := t
			for len(st) > 0 {
				temp.Next = st[0]
				st = st[1:]
				temp = temp.Next
			}
			temp.Next = nil
			if root.Val == -1 {
				root = t
				p = root
			} else {
				for p.Next != nil {
					p = p.Next
				}
				p.Next = t
			}
		}
	}
	return root
}

//--------------------------non recursive looks bad--------------------

//--------------------------better approach--------------------
func reverse_nodes_in_k_groups_better(head *ListNode, k int) *ListNode {
	ptr := head
	var tail *ListNode
	var new_head *ListNode
	for ptr != nil {
		cnt := 0
		ptr = head
		for cnt < k && ptr != nil {
			ptr = ptr.Next
			cnt++
		}
		if cnt == k {
			rev_head := reverse(head, k)
			if new_head == nil {
				new_head = rev_head
			}
			if tail != nil {
				tail.Next = rev_head
			}
			tail = head
			head = ptr
		}
	}
	if tail != nil {
		tail.Next = head
	}
	if new_head == nil {
		return head
	}
	return new_head
}

func reverse(head *ListNode, k int) *ListNode {
	var new_head *ListNode
	ptr := head
	for k > 0 {
		next_node := ptr.Next
		ptr.Next = new_head
		new_head = ptr
		ptr = next_node
		k--
	}
	return new_head
}

//--------------------------better approach--------------------

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(4)
	head.Next.Next.Next.Next = New(5)
	fmt.Println(reverse_nodes_in_k_groups_better(head, 2))
}
