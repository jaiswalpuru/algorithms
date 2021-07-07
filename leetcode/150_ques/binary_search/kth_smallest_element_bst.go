/*Given the root of a binary search tree, and an integer k,
return the kth (1-indexed) smallest element in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1
*/

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func New(data int) *Node { return &Node{data: data, left: nil, right: nil} }

func _kth_smallest(root *Node, arr *[]int) {
	if root == nil {
		return
	}
	_kth_smallest(root.left, arr)
	*arr = append(*arr, root.data)
	_kth_smallest(root.right, arr)
}

func kth_smallest(root *Node, k int) int {
	arr := []int{}
	_kth_smallest(root, &arr)
	return arr[k-1]
}

func main() {
	root := New(3)
	root.left = New(1)
	root.right = New(4)
	root.left.right = New(2)
	fmt.Println(kth_smallest(root, 1))
	root = New(5)
	root.left = New(3)
	root.right = New(6)
	root.left.left = New(2)
	root.left.right = New(4)
	root.left.left.left = New(1)
	fmt.Println(kth_smallest(root, 3))
}
