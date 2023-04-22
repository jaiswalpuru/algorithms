package main

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head       *ListNode
	indexValue map[int]int
	size       int
}

func Constructor(head *ListNode) Solution {
	indexMap := updateIndex(head)
	size := getLength(head)
	return Solution{head: head, indexValue: indexMap, size: size}
}

func getLength(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + getLength(head.Next)
}

func updateIndex(head *ListNode) map[int]int {
	index := 0
	indexMap := make(map[int]int)

	for head != nil {
		indexMap[index] = head.Val
		head = head.Next
		index++
	}

	return indexMap
}

func (this *Solution) GetRandom() int {
	return this.indexValue[rand.Intn(this.size)]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
