package main

import (
	"container/heap"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Pair struct {
	val int
	dis float64
}

type M []Pair

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i].dis < m[j].dis }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val} }

func closest_binary_search_tree_value_two(root *TreeNode, target float64, k int) []int {
	mh := &M{}
	inorder(root, target, mh)
	res := []int{}
	for k > 0 {
		res = append(res, heap.Pop(mh).(Pair).val)
		k--
	}
	return res
}

func inorder(root *TreeNode, target float64, mh *M) {
	if root == nil {
		return
	}
	inorder(root.Left, target, mh)
	heap.Push(mh, Pair{val: root.Val, dis: abs(float64(root.Val) - target)})
	inorder(root.Right, target, mh)
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -1 * a
}

//------------------using quickselect---------------------
func closest_binary_search_tree_value_two_quick(root *TreeNode, target float64, k int) []int {
	res := []int{}
	_inorder(root, &res)
	quickselect(0, len(res)-1, res, k, target)
	return res[:k]
}

func partition(l, r int, pivot int, nums []int, target float64) int {
	dis_piv := abs(float64(nums[pivot]) - target)
	nums[pivot], nums[r] = nums[r], nums[pivot]
	store_ind := l
	for i := l; i <= r; i++ {
		if abs(float64(nums[i])-target) < dis_piv {
			nums[store_ind], nums[i] = nums[i], nums[store_ind]
			store_ind++
		}
	}
	nums[store_ind], nums[r] = nums[r], nums[store_ind]
	return store_ind
}

func quickselect(l, r int, nums []int, k int, target float64) {
	if l >= r {
		return
	}
	pivot := l
	pivot = partition(l, r, pivot, nums, target)
	if k == pivot {
		return
	} else if k < pivot {
		quickselect(l, pivot-1, nums, k, target)
	} else {
		quickselect(pivot+1, r, nums, k, target)
	}
}

func _inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	_inorder(root.Left, res)
	*res = append(*res, root.Val)
	_inorder(root.Right, res)
}

//------------------using quickselect---------------------

func main() {
	root := New(4)
	root.Left = New(2)
	root.Right = New(5)
	root.Left.Left = New(1)
	root.Left.Right = New(3)
	fmt.Println(closest_binary_search_tree_value_two_quick(root, 3.714286, 2))
}
