/*
	Given array of integers, return a new array where each element in the new array is the number of smaller
	elements to the right of that element in original array
	eg: arr := [3,4,9,6,1] return [1,1,2,1,0]

	Brute force : Create a new array of the same length and for each element in the original array count the number
	of smaller elements to the right. O(n^2)
*/

package main

import "fmt"

type Node struct {
	data        int
	count       int
	left, right *Node
}

func New(data, count int) *Node { return &Node{data: data, count: count, left: nil, right: nil} }

func insert(root **Node, val, cnt int) int {
	if (*root) == nil {
		(*root) = New(val, 0)
		return cnt
	}

	if (*root).data < val {
		return ((*root).count) + insert(&(*root).right, val, cnt+1)
	} else {
		(*root).count++
		return insert(&(*root).left, val, cnt)
	}
}

func main() {

	arr := []int{12, 1, 2, 3, 0, 11, 4}
	result := make([]int, len(arr))
	var root *Node
	for i := len(arr) - 1; i >= 0; i-- {
		result[i] = insert(&root, arr[i], 0)
	}
	fmt.Println(result)
}
