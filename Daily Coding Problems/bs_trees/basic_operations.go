package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data        int
	left, right *Node
}

func New(data int) *Node { return &Node{data: data} }

func insert(root *Node, data int) {
	if root == nil {
		root = New(data)
	} else {
		_insert(root, data)
	}
}

func _insert(node *Node, data int) {
	if data < node.data {
		if node.left == nil {
			node.left = New(data)
		} else {
			insert(node.left, data)
		}
	} else {
		if node.right == nil {
			node.right = New(data)
		} else {
			insert(node.right, data)
		}
	}
}

func inorder(root *Node, str *string) {
	if root == nil {
		return
	}
	inorder(root.left, str)
	*str += strconv.Itoa(root.data) + "->"
	inorder(root.right, str)
}

func (node *Node) String() string {
	str := ""
	inorder(node, &str)
	return fmt.Sprintf("%v", str)
}

func find(root *Node, data int) bool {
	if root == nil {
		return false
	} else {
		return _find(root, data)
	}
}

func _find(node *Node, data int) bool {
	if node == nil {
		return false
	} else if node.data == data {
		return true
	} else if node.data > data {
		return _find(node.left, data)
	} else {
		return _find(node.right, data)
	}
}

func main() {
	root := New(7)
	insert(root, 5)
	insert(root, -1)
	insert(root, 6)
	insert(root, 10)
	insert(root, 25)
	fmt.Println(root)
	fmt.Println(find(root, 10))
	fmt.Println(find(root, 30))
}
