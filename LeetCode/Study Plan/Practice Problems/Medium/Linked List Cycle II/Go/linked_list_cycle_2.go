package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func linked_list_cycle_two(root *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	hash_map := make(map[*ListNode]bool)
	hash_map[head] = true
	for head != nil {
		if hash_map[head.Next] {
			return head.Next
		}
		hash_map[head.Next] = true
		head = head.Next
	}
	return nil
}

func get_intersection_node(head *ListNode) *ListNode {
	sp, fp := head, head

	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
		if sp == fp {
			return fp
		}
	}
	return nil
}

func linked_list_cycle_two_sp_fp(root *ListNode) *ListNode {
	v := get_intersection_node(head)
	if v == nil {
		return nil
	}
	p1, p2 := head, v
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

func main() {
	head := New(3)
	head.Next = New(2)
	head.Next.Next = New(0)
	head.Next.Next.Next = New(-4)
	head.Next.Next.Next.Next = head.Next
	fmt.Println(linked_list_cycle_two(head))
}
