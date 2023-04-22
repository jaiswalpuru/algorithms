package main

type BSTIterator struct {
	j     int
	inord []int
	root  *TreeNode
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
	inorder(root, &order)
	return BSTIterator{j: -1, root: root, inord: order}
}

func (this *BSTIterator) HasNext() bool {
	return this.j < len(this.inord)-1
}

func (this *BSTIterator) Next() int {
	val := this.inord[this.j+1]
	this.j += 1
	return val
}

func (this *BSTIterator) HasPrev() bool {
	return this.j > 0
}

func (this *BSTIterator) Prev() int {
	val := this.inord[this.j-1]
	this.j -= 1
	return val
}

func main() {

}
