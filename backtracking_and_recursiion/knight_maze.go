package main

import "fmt"

//Given a square board of NXN size, the position of knight and position of the target is given.
//We need to find the minimum number of steps the knight will require to reach the target position

//to check for all possible postions around the knight
var (
	row = []int{2, 2, -2, -2, 1, 1, -1, -1}
	col = []int{-1, 1, 1, -1, 2, -2, 2, -2}
)

type Node struct {
	x, y, dist int
}

func newNode(x, y, start int) Node {
	return Node{x: x, y: y, dist: start}
}

func validPosition(x, y, N int) bool {
	if x < 0 || y < 0 || x >= N || y >= N {
		return false
	}
	return true
}

func main() {

	var N int // length of the board
	//fmt.Scan(&N)
	N = 8
	visited := make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, N)
	}

	startPos := newNode(0, 7, 0)
	dstPos := newNode(7, 0, 0)
	q := []Node{}
	q = append(q, startPos)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.x == dstPos.x && curr.y == dstPos.y {
			fmt.Println("Destination reached : ", curr.dist)
			return
		}
		for i := 0; i < 8; i++ {
			x := curr.x + row[i]
			y := curr.y + col[i]
			if validPosition(x, y, N) && !visited[x][y] {
				visited[x][y] = true
				q = append(q, newNode(x, y, curr.dist+1))
			}
		}
	}
}
