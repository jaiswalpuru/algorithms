package main

import "fmt"

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

func subtree_sum_cnt_nodes(root *TreeNode, sum *int, cnt *int) {
	if root == nil {
		return
	}
	subtree_sum_cnt_nodes(root.Left, sum, cnt)
	subtree_sum_cnt_nodes(root.Right, sum, cnt)
	*sum += root.Val
	*cnt += 1
}

func count_node_equal_to_average_of_subtree(root *TreeNode) int {
	res := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.Left != nil {
			q = append(q, curr.Left)
		}
		if curr.Right != nil {
			q = append(q, curr.Right)
		}

		sum, cnt := 0, 0
		subtree_sum_cnt_nodes(curr, &sum, &cnt)
		if sum/cnt == curr.Val {
			res++
		}
	}
	return res
}

func main() {
	root := New(4)
	root.Left = New(8)
	root.Right = New(5)
	root.Right.Right = New(6)
	root.Left.Left = New(0)
	root.Left.Right = New(1)
	fmt.Println(count_node_equal_to_average_of_subtree(root))
}
