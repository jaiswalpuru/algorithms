package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func New(val int) *TreeNode { return &TreeNode{Val: val, Left: nil, Right: nil} }

type CBTInserter struct{ root *TreeNode }

func Constructor(root *TreeNode) CBTInserter { return CBTInserter{root} }

func (this *CBTInserter) Insert(val int) int {
	q := []*TreeNode{this.root}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.Left == nil {
			curr.Left = New(val)
			return curr.Val
		}
		if curr.Right == nil {
			curr.Right = New(val)
			return curr.Val
		}

		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}
	return -1
}

func (this *CBTInserter) Get_root() *TreeNode { return this.root }

func main() {}
