package main

type Node struct {
	Val int
	Children []*Node
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {}

func new_treenode(val int) *TreeNode {return &TreeNode{Val:val, Left:nil, Right:nil}}

func new_node(val int) *Node {return &Node{Val: val, Children:make([]*Node, 0)}}

func Constructor() *Codec {return new(Codec)}

func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}

	_root := new_treenode(root.Val)
	if len(root.Children) > 0{
		_root.Left = this.encode(root.Children[0])
	}

	right := _root.Left
	for i:=1;i<len(root.Children);i++{
		right.Right = this.encode(root.Children[i])
		right = right.Right
	}
	return _root
}

func (this *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	_root := new_node(root.Val)
	child := root.Left
	for child != nil {
		_root.Children = append(_root.Children, this.decode(child))
		child = child.Right
	}
	return _root
}

func main() {

}