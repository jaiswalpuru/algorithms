package main

type ListNode struct {
	Val int
	Next *ListNode
}

func New(val int)*ListNode { return &ListNode{Val : val, Next nil}}

//-------------------------------recursive approach--------------------------------------
func reverse_list_recursive(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	root :=  _recursive_reverse(head)
	head.Next = nil
	return root
}

func _recursive_reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	temp := _recursive_reverse(head.Next)
	if temp == nil {
		return head
	}
	head.Next.Next = head
	return temp
}
//-------------------------------recursive approach--------------------------------------

func main() {
	head := New(1)
	head.Next = New(2)
	head.Next.Next = New(3)
	head.Next.Next.Next = New(4)
	fmt.Println(reverse_list_recursive(head))
}
