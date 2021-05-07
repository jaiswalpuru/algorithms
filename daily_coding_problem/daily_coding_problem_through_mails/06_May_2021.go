/*
You are given an M by N matrix consisting of booleans that represents a board. Each True boolean represents a wall.
Each False boolean represents a tile you can walk on.

Given this matrix, a start coordinate, and an end coordinate,
return the minimum number of steps required to reach the end coordinate from the start.
If there is no possible path, then return null. You can move up, left, down, and right.
You cannot move through walls. You cannot wrap around the edges of the board.

For example, given the following board:

[[f, f, f, f],
[t, t, f, t],
[f, f, f, f],
[f, f, f, f]]
and start = (3, 0) (bottom left) and end = (0, 0) (top left), the minimum number of steps required
to reach the end is 7, since we would need to go through (1, 2) because there is a wall everywhere
else on the second row.
*/

/*
This can be solved using Lee algorithm
1) We start with the source cell and call BFS
2) Maintain a queue and initialize with the source cell
3) Maintain a boolean array of the size of the board and initialize all values to false
	PUSH(src_coord)
	REPEAT :
		DEQUEUE()
		RETURN IF DEST REACHED
		FOR EACH Adjecent cell if value is 0, and not yet visited ENQUEUE and mark them as visited
	UNTIL : QUEUE NOT EMPTY
*/

package main

import "fmt"

type Pair struct {
	a, b int
}

type Triplet struct {
	p     Pair
	count int
}

var (
	cnt   = 0
	board = [][]bool{
		{f, f, f, f},
		{t, t, f, t},
		{f, f, f, f},
		{f, f, f, f},
	}
	row = []int{-1, 0, 0, 1}
	col = []int{0, -1, 1, 0}
)

const (
	f = false
	t = true
)

func isValid(r, c int) bool { return r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) }

func findShortestPath(s, e Pair) int {
	if board[s.a][s.b] || board[e.a][e.b] {
		return -1
	}

	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	visited[s.a][s.b] = true

	queue := []Triplet{}
	queue = append(queue, Triplet{p: s, count: 0})

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.p.a == e.a && curr.p.b == e.b {
			return curr.count
		}

		for i := 0; i < 4; i++ {
			r := curr.p.a + row[i]
			c := curr.p.b + col[i]
			if isValid(r, c) && !board[r][c] && !visited[r][c] {
				visited[r][c] = true
				fmt.Println("r,c : ", r, c)
				queue = append(queue, Triplet{Pair{r, c}, curr.count + 1})
			}
		}
	}
	return -1
}

func main() {
	s := Pair{3, 0}
	e := Pair{0, 0}
	fmt.Println(findShortestPath(s, e))
}
