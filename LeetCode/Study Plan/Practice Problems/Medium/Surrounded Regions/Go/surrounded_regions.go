package main

import "fmt"

func surround_regions(board [][]byte) [][]byte {
	n, m := len(board), len(board[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || j == 0 || i == n-1 || j == m-1) && board[i][j] == 'o' {
				dfs(&board, i, j, n, m)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'o' {
				board[i][j] = 'x'
			}
			if board[i][j] == '#' {
				board[i][j] = 'o'
			}
		}
	}

	return board
}

func dfs(board *[][]byte, i, j, n, m int) {
	if i < 0 || j < 0 || i >= n || j >= m || (*board)[i][j] == '#' || (*board)[i][j] == 'x' {
		return
	}
	(*board)[i][j] = '#'
	dfs(board, i+1, j, n, m)
	dfs(board, i-1, j, n, m)
	dfs(board, i, j+1, n, m)
	dfs(board, i, j-1, n, m)
}

func main() {
	fmt.Println(surround_regions([][]byte{
		{'x', 'x', 'x', 'x'},
		{'x', 'o', 'o', 'x'},
		{'x', 'x', 'o', 'x'},
		{'x', 'o', 'x', 'x'},
	}))
}
