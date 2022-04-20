package main

type TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

type BSTIterator struct {
	i     int
	n     int
	order []int
}

func inorder(root *TreeNode, order *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, order)
	*order = append(*order, root.Val)
	inorder(root.Right, order)
}

func Constructor(root *TreeNode) BSTIterator {
	order := []int{}
	i := 0
	inorder(root, &order)
	n := len(order)
	return BSTIterator{i: i, n: n, order: order}
}

func (this *BSTIterator) Next() int {
	val := this.order[this.i]
	this.i += 1
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.i < this.n
}

func main() {}
