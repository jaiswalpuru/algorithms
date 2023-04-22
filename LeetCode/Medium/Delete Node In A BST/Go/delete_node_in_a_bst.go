package main

type  TreeNode struct {
	Val int
	Left,
	Right *TreeNode
}

func New(val int) *TreeNode {&TreeNode{Val:val, Left:nil, Right:nil}}

func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func delete_node_in_bst(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = delete_node_in_bst(root.Right, key)
	} else if key < root.Val {
		root.Left = delete_node_in_bst(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = successor(root)
			root.Right = delete_node_in_bst(root.Right, root.Val)
		} else {
			root.Val = predecessor(root)
			root.Left = delete_node_in_bst(root.Left, root.Val)
		}
	}

	return root
}

func main(){
	root := New(5)
	root.Left = New(3)
	root.Right = New(6)
	root.Right.Right = New(7)
	root.Left.Left = New(2)
	root.Left.Right = New(4)
	fmt.Println(delete_node_in_bst(root, 3))
}