package main

import "fmt"

func snakesAndLadders(board [][]int) int {
	n := len(board)
	boardFull := make([]int, n*n+1)
	k := 1
	for i := 0; i < n/2; i++ {
		board[i], board[n-i-1] = board[n-i-1], board[i]
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				boardFull[k] = board[i][j]
				k++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				boardFull[k] = board[i][j]
				k++
			}
		}
	}
	q := []int{}
	if boardFull[1] != -1 {
		q = append(q, boardFull[1])
	} else {
		q = append(q, 1)
	}
	dis := 0
	visited := make(map[int]bool)
	for len(q) > 0 {
		qSize := len(q)
		for j := 0; j < qSize; j++ {
			curr := q[0]
			q = q[1:]
			if curr == n*n {
				return dis
			}
			for i := curr + 1; i <= min(curr+6, n*n); i++ {
				nextPos := i
				if boardFull[i] != -1 {
					nextPos = boardFull[i]
				}
				if !visited[nextPos] {
					visited[nextPos] = true
					q = append(q, nextPos)
				}
			}
		}
		dis++
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(snakesAndLadders([][]int{
		{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1},
	}))
}
