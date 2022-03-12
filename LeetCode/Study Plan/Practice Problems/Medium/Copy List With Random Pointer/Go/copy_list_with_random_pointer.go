package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func New(val int) *Node { return &Node{Val: val, Next: nil, Random: nil} }

//------------------------------------using iterative----------------------------------------------
func get_clone(node *Node, visited *map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := (*visited)[node]; !ok {
		(*visited)[node] = New(node.Val)
	}
	return (*visited)[node]
}

func deep_copy(head *Node) *Node {
	if head == nil {
		return nil
	}

	old_node := head

	new_node := New(old_node.Val)
	visited := make(map[*Node]*Node)
	visited[old_node] = new_node

	for old_node != nil {
		new_node.Random = get_clone(old_node.Random, &visited)
		new_node.Next = get_clone(old_node.Next, &visited)
		old_node = old_node.Next
		new_node = new_node.Next
	}

	return visited[head]
}

//------------------------------------using iterative----------------------------------------------

//----------------------------------------using recursion----------------------------------------
func deep_copy_recursion(root *Node) *Node {
	hash_map := make(map[*Node]*Node)
	return _deep_copy_recursion(root, &hash_map)
}

func _deep_copy_recursion(root *Node, hash_map *map[*Node]*Node) *Node {
	if root == nil {
		return nil
	}
	if _, ok := (*hash_map)[root]; ok {
		return (*hash_map)[root]
	}

	new_node := New(root.Val)
	(*hash_map)[root] = new_node

	new_node.Next = _deep_copy_recursion(root.Next, hash_map)
	new_node.Random = _deep_copy_recursion(root.Random, hash_map)
	return new_node
}

//--------------------------------------------------------------------------------------------------

func main() {
	head := New(7)
	head.Next = New(13)
	head.Next.Random = head
	head.Next.Next = New(11)
	head.Next.Next.Next = New(10)
	head.Next.Next.Next.Next = New(1)
	head.Next.Next.Random = head.Next.Next.Next.Next
	head.Next.Next.Next.Random = head.Next.Next
	head.Next.Next.Next.Next.Random = head
	fmt.Println(deep_copy_recursion(head))
}
