package main

import "fmt"

func is_safe(board [][]int, row, col int) bool {
	//check current row on left side
	for i := 0; i < col; i++ {
		if board[row][i] == 1 {
			return false
		}
	}

	//check upper diagonal on left
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	//check lower diagonal on left
	for i, j := row, col; j >= 0 && i < len(board); i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}

func n_queens(board [][]int, col int) bool {
	//if all the queens are placed
	if col >= len(board) {
		return true
	}

	//consider this col and try placing this queen to every row one by one
	for i := 0; i < len(board); i++ {
		if is_safe(board, i, col) {
			board[i][col] = 1
			//call n_queens for the rest of the queens
			if n_queens(board, col+1) {
				return true
			}
			board[i][col] = 0
		}
	}
	//queen cannot be placed in any location
	return false
}

func main() {
	board := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	t := n_queens(board, 0)
	if !t {
		fmt.Println("solution does not exist")
	} else {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				fmt.Print(board[i][j], " ")
			}
			fmt.Println()
		}
	}
}
