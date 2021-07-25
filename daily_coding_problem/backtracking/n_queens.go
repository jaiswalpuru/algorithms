//N Queens

package main

import "fmt"

func is_safe(board [][]int, r, c int) bool {
	//check left side of the board
	for i := 0; i < c; i++ {
		if board[r][i] == 1 {
			return false
		}
	}

	//check upper diagonal on left
	for i, j := r, c; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	n := len(board)
	//check lower left diagonal
	for i, j := r, c; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func _n_queens(board *[][]int, col int) bool {
	if col >= len(*board) {
		return true
	}

	for i := 0; i < len(*board); i++ {
		if is_safe(*board, i, col) {
			(*board)[i][col] = 1
			if _n_queens(board, col+1) {
				return true
			}
			(*board)[i][col] = 0
		}
	}
	return false
}

func n_queens(n int) {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	ans := _n_queens(&board, 0)

	if !ans {
		fmt.Println("Solution doesn't exist!")
	} else {
		for i := 0; i < n; i++ {
			fmt.Println(board[i])
		}
	}
}

func main() {
	n_queens(4)
}
