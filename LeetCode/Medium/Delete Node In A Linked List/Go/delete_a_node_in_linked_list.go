package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val} }

//-------------------delete a node in linked list---------------- O(n)
func delete_a_node_in_linked_list(node *ListNode) {
	prev := node
	for node.Next != nil {
		prev = node
		node.Val, node.Next.Val = node.Next.Val, node.Val
		node = node.Next
	}
	prev.Next = nil
}

//-------------------delete a node in linked list----------------

//-------------------delete a node in linked list---------------- O(1)
func delete_a_node_in_linked_list_eff(node *ListNode) {
	if node.Next == nil {
		node = nil
	}
	node.Next.Val, node.Val = node.Val, node.Next.Val
	node.Next = node.Next.Next
}

//-------------------delete a node in linked list----------------

func main() {
	head := New(4)
	head.Next = New(5)
	head.Next.Next = New(1)
	head.Next.Next.Next = New(9)
	delete_a_node_in_linked_list(head.Next)
}
