package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(val int) *ListNode { return &ListNode{Val: val, Next: nil} }

func linked_list_components(head *ListNode, nums []int) int {
	hash_map := make(map[int]int)
	temp := head
	for temp != nil {
		if temp.Next != nil {
			hash_map[temp.Val] = temp.Next.Val
		} else {
			hash_map[temp.Val] = -1
		}
		temp = temp.Next
	}

	v := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v[nums[i]] = -1
	}

	for i := 0; i < len(nums); i++ {
		val := hash_map[nums[i]]
		if _, ok := v[val]; ok {
			v[nums[i]] = val
		}
	}

	cnt := 0
	for _, val := range v {
		if val == -1 {
			cnt++
		}
	}

	return cnt
}

func main() {
	root := New(0)
	root.Next = New(1)
	root.Next.Next = New(2)
	root.Next.Next.Next = New(3)
	fmt.Println(linked_list_components(root, []int{0, 1, 3}))
}
