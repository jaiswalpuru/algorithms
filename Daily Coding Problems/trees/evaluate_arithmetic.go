/*
	Expression trees -> all the leaves are integers and non leaves are characters
	Hint : Post Order traversal
*/

package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

const (
	plus   = int('+')
	minus  = int('-')
	mult   = int('*')
	divide = int('/')
)

func New(data int) *Node { return &Node{data: data} }

func evaluate(root *Node) int {
	if root.data == plus {
		return evaluate(root.left) + evaluate(root.right)
	} else if root.data == minus {
		return evaluate(root.left) - evaluate(root.right)
	} else if root.data == mult {
		return evaluate(root.left) * evaluate(root.right)
	} else if root.data == divide {
		return evaluate(root.left) / evaluate(root.right)
	} else {
		return root.data
	}
}

func main() {
	root := New(mult)
	root.left = New(plus)
	root.right = New(plus)
	root.left.left = New(3)
	root.left.right = New(2)
	root.right.left = New(4)
	root.right.right = New(5)
	fmt.Println(evaluate(root))
}
