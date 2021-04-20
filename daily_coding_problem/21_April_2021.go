/*
A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.

Given the root to a binary tree, count the number of unival subtrees.

For example, the following tree has 5 unival subtrees:

   0
  / \
 1   0
   / \
  1   0
 / \
1   1

*/

package main

import "fmt"

type TreeNode struct {
	data        interface{}
	left, right *TreeNode
}

func isUnival(root *TreeNode) bool {
	if root == nil {
		return true
	}

	//check if data parent and child matches
	if root.left != nil && root.data != root.left.data {
		return false
	}

	if root.right != nil && root.data != root.right.data {
		return false
	}

	if isUnival(root.left) && isUnival(root.right) {
		return true
	}
	return false
}

func countUnival(root *TreeNode) int { //O(n^2) soln
	if root == nil {
		return 0
	}
	totalCount := countUnival(root.left) + countUnival(root.right)
	if isUnival(root) {
		totalCount += 1
	}
	return totalCount
}

func main() {
	root := &TreeNode{data: 0, left: nil, right: nil}
	root.left = &TreeNode{data: 1, left: nil, right: nil}
	root.right = &TreeNode{data: 0, left: nil, right: nil}
	root.right.right = &TreeNode{data: 0, left: nil, right: nil}
	root.right.left = &TreeNode{data: 1, left: nil, right: nil}
	root.right.left.left = &TreeNode{data: 1, left: nil, right: nil}
	root.right.left.right = &TreeNode{data: 1, left: nil, right: nil}
	fmt.Println(countUnival(root))
	totalCount, _ := helper(root)
	fmt.Println(totalCount)
}

func helper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftCount, isLeftUnival := helper(root.left)
	rightCount, isRightUinval := helper(root.right)
	is_Unival := true
	if !isLeftUnival || !isRightUinval {
		is_Unival = false
	}
	if root.left != nil && root.left.data != root.data {
		is_Unival = false
	}

	if root.right != nil && root.right.data != root.data {
		is_Unival = false
	}

	if is_Unival {
		return leftCount + rightCount + 1, true
	} else {
		return leftCount + rightCount, false
	}
}
