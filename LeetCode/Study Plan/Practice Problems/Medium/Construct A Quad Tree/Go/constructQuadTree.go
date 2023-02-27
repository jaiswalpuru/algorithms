package main

import "fmt"

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func New(val, isLeaf bool) *Node { return &Node{val, isLeaf, nil, nil, nil, nil} }

func construct(grid [][]int) *Node {
	size := len(grid)
	return recur(0, 0, grid, size)
}

func recur(x, y int, grid [][]int, size int) *Node {
	if isSameValue(grid, x, y, size) {
		return New(grid[x][y] == 1, true)
	} else {
		//call recur over the four submatrices
		root := New(false, false)
		root.TopLeft = recur(x, y, grid, size/2)
		root.TopRight = recur(x, y+size/2, grid, size/2)
		root.BottomLeft = recur(x+size/2, y, grid, size/2)
		root.BottomRight = recur(x+size/2, y+size/2, grid, size/2)
		return root
	}
}

func isSameValue(grid [][]int, x, y, size int) bool {
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if grid[x][y] != grid[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(construct([][]int{
		{0, 1}, {1, 0},
	}))
}
