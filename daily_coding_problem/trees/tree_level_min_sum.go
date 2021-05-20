package main

import "fmt"

type Node struct {
	data        int
	left, right *Node
}

type Pair struct {
	node  *Node
	level int
}

func New(data int) *Node { return &Node{data: data} }

func min_level_sum(root *Node) int {
	if root == nil {
		return 0
	}
	queue := []Pair{}
	queue = append(queue, Pair{node: root, level: 0})
	level_sum := make(map[int]int)

	for len(queue) > 0 {
		curr, level := queue[0].node, queue[0].level
		queue = queue[1:]
		level_sum[level] += curr.data
		if curr.left != nil {
			queue = append(queue, Pair{node: curr.left, level: level + 1})
		}
		if curr.right != nil {
			queue = append(queue, Pair{node: curr.right, level: level + 1})
		}
	}

	min_level := 0
	fmt.Println(level_sum)
	for k, v := range level_sum {
		if level_sum[min_level] > v {
			min_level = k
		}
	}
	return min_level
}

func main() {
	root := New(1)
	root.left = New(2)
	root.right = New(3)
	root.right.left = New(4)
	root.right.right = New(5)
	fmt.Println(min_level_sum(root))
}
