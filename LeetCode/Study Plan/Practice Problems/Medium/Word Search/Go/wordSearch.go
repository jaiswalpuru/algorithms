package main

import "fmt"

var dir = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if word[0] == board[i][j] {
				if backtrack(board, i, j, 0, word, &visited) {
					return true
				}
			}
		}
	}
	return false
}

func isSafe(i, j, n, m int, visited [][]bool) bool {
	return i >= 0 && j >= 0 && i < n && j < m && !visited[i][j]
}

func backtrack(board [][]byte, i, j, k int, word string, visited *[][]bool) bool {
	if k == len(word) {
		return true
	}
	if !isSafe(i, j, len(board), len(board[0]), *visited) || word[k] != board[i][j] {
		return false
	}
	(*visited)[i][j] = true
	for d := 0; d < 4; d++ {
		x, y := dir[d][0]+i, dir[d][1]+j
		if backtrack(board, x, y, k+1, word, visited) {
			return true
		}
	}
	(*visited)[i][j] = false
	return false
}

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE"))
}
