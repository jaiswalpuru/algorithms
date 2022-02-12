package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

//------------------using bfs-----------------------
func deepest_leave_sum(root *TreeNode) int {
	q := []*TreeNode{root}
	sum := 0
	for len(q) > 0 {
		n := len(q)
		sum = 0
		for i := 0; i < n; i++ {
			sum += q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
	}
	return sum
}

//------------------using bfs-----------------------

//------------------using dfs-----------------------
func deepest_leave_sum_dfs(root *TreeNode) int {
	type Pair struct {
		val *TreeNode
		lvl int
	}
	stack := []Pair{{root, 0}}
	curr_depth := 0
	depth := 0
	sum := 0
	for len(stack) > 0 {
		n := len(stack) - 1
		curr := stack[n].val
		curr_depth = stack[n].lvl

		if depth < curr_depth {
			depth = curr_depth
			sum = curr.Val
		} else if depth == curr_depth {
			sum += curr.Val
		}
		stack = stack[:n]
		if curr.Left != nil {
			stack = append(stack, Pair{curr.Left, curr_depth + 1})
		}
		if curr.Right != nil {
			stack = append(stack, Pair{curr.Right, curr_depth + 1})
		}
	}
	return sum
}

//------------------using dfs-----------------------

func main() {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Left.Left.Left = New(7)
	root.Right.Right = New(6)
	root.Right.Right.Right = New(8)
	fmt.Println(deepest_leave_sum_dfs(root))
}
