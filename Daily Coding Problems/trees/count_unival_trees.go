/*
	Unival tree is a tree in which all the nodes under it have same value.
	   0
  	  / \
     1   0
        / \
       1   0
      / \
     1   1
*/

package main

import (
	"fmt"
)

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data} }

func is_unival(root *Node) bool { return unival_helper(root, root.data) }

func unival_helper(root *Node, val int) bool {
	if root == nil {
		return true
	}
	if root.data == val {
		return unival_helper(root.left, val) && unival_helper(root.right, val)
	}
	return false
}

//takes O(n^2)
func count_unival_trees_recursive_non_efficient(root *Node) int {
	if root == nil {
		return 0
	}
	left := count_unival_trees_recursive_non_efficient(root.left)
	right := count_unival_trees_recursive_non_efficient(root.right)

	if is_unival(root) {
		return left + right + 1
	} else {
		return left + right
	}
}

//more efficient recursive approach
func count_unival_trees_recursive_efficient(root *Node) (int, bool) {
	if root == nil {
		return 0, true
	}

	left_count, is_left_unival := count_unival_trees_recursive_efficient(root.left)
	right_count, is_right_unival := count_unival_trees_recursive_efficient(root.right)
	total := left_count + right_count

	if is_left_unival && is_right_unival {
		if root.left != nil && root.data != root.left.data {
			return total, false
		}
		if root.right != nil && root.data != root.right.data {
			return total, false
		}
		return total + 1, true
	}
	return total, false
}

//non-recursive approach
func count_unival_trees(root *Node) int {
	cnt := 0
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if (curr.left == curr.right) || (curr.left.data == curr.right.data) {
			cnt++
		}
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
	return cnt
}

func main() {
	root := New(0)
	root.left = New(1)
	root.right = New(0)
	root.right.left = New(1)
	root.right.left.left = New(1)
	root.right.left.right = New(1)
	root.right.right = New(0)

	fmt.Println(count_unival_trees(root)) //much efficient

	fmt.Println(count_unival_trees_recursive_non_efficient(root))

	cnt, _ := count_unival_trees_recursive_efficient(root)
	fmt.Println(cnt)
}
