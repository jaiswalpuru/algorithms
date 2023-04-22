package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func New(val int) *Node { return &Node{Val: val, Children: nil} }

//clone graph solution wont work here as numnber of nodes are 10^4 so while doing a bfs traverse it will run out
// of memoory

func clone_n_ary_tree(root *Node) *Node {
	return _clone_recursion(root)
}

func _clone_recursion(root *Node) *Node {
	if root == nil {
		return nil
	}
	node_copy := New(root.Val)
	for i := 0; i < len(root.Children); i++ {
		node_copy.Children = append(node_copy.Children, _clone_recursion(root.Children[i]))
	}
	return node_copy
}

//-----------dfs iterative-----------------------------------------------
func _clone_dfs(root *Node) *Node {
	if root == nil {
		return nil
	}
	new_node := New(root.Val)
	type P struct{ old_node, new_node *Node }
	st := []P{{root, new_node}} //push the original node and cloned node
	for len(st) > 0 {
		curr := st[len(st)-1]
		old_node := curr.old_node
		new_node := curr.new_node
		for i := 0; i < len(old_node.Children); i++ {
			child_new_node := New(old_node.Children[i].Val)
			new_node.Children = append(new_node.Children, child_new_node)
			st = append(st, P{old_node: old_node.Children[i], new_node: child_new_node})
		}
	}
	return new_node
}

//---------------------------------------------------------------------

func _clone_bfs(root *Node) *Node {
	if root == nil {
		return nil
	}
	type P struct{ old_node, new_node *Node }
	new_node := New(root.Val)
	q := []P{{root, new_node}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		o, n := curr.old_node, curr.new_node
		for i := 0; i < len(o.Children); i++ {
			child_new_node := New(o.Children[i].Val)
			n.Children = append(n.Children, child_new_node)
			q = append(q, P{o.Children[i], child_new_node})
		}
	}
	return new_node
}

func main() {
	root := New(1)
	root.Children = append(root.Children, New(3))
	root.Children = append(root.Children, New(2))
	root.Children = append(root.Children, New(4))
	root.Children[0].Children = append(root.Children[0].Children, New(5))
	root.Children[0].Children = append(root.Children[0].Children, New(6))
	fmt.Println(clone_n_ary_tree(root))
}
