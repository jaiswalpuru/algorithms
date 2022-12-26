package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func New(val int) *Node { return &Node{Val: val, Next: nil, Random: nil} }

//------------------------------------using iterative----------------------------------------------
func get_clone(node *Node, mp *map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := (*mp)[node]; !ok {
		(*mp)[node] = New(node.Val)
	}
	return (*mp)[node]
}

func deep_copy(head *Node) *Node {
	if head == nil {
		return nil
	}
	temp := head
	mp := make(map[*Node]*Node)
	new_node := New(temp.Val)
	mp[temp] = new_node
	for temp != nil {
		new_node.Random = get_clone(temp.Random, &mp)
		new_node.Next = get_clone(temp.Next, &mp)
		temp = temp.Next
		new_node = new_node.Next
	}
	return mp[head]
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
