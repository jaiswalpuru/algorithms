package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New() *TreeNode { return &TreeNode{Val: 0, Left: nil, Right: nil} }

var mp = make(map[int][]*TreeNode)

func FBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}

	if _, ok := mp[n]; !ok {
		ans := []*TreeNode{}
		if n == 1 {
			ans = append(ans, New())
		} else {
			for p := 0; p < n; p++ {
				q := n - 1 - p
				for _, left := range FBT(p) {
					for _, right := range FBT(q) {
						t := New()
						t.Left = left
						t.Right = right
						ans = append(ans, t)
					}
				}

			}
		}
		mp[n] = append(mp[n], ans...)
	}
	return mp[n]
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(FBT(7))
}

//Not relevant to this ques, just for information
/*
	In a "full" tree, there are an odd number of nodes and every second node in order is a leaf.
	If you remove all these leaves, you are left with a binary tree that might not be full. For any
	(maybe not full) binary tree, there is exactly one way to add a leaf at the start, the end, and between each pair of
	nodes, to make a full binary tree.
	So there is a 1-1 correspondence between binary trees with n nodes, and full trees with 2n+1 codes.
	C(n) -- the catalan number -- is the number of binary trees with n nodes, and also therefore the number
	of "full" trees with 2n+1 nodes.
	The number of full binary trees with n nodes is therefore C((n-1)/2).
	Since you can't have half a node, the answer is 0 when n is even.
	Reference : stackoverflow

*/
