/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
The '.' character indicates empty cells.
*/

package main

import "fmt"

func can_place(board [][]byte, row, col int, char byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == char {
			return false
		}

		if board[row][i] == char {
			return false
		}

		if board[3*(row/3)+i/3][3*(col/3)+i%3] == char {
			return false
		}
	}
	return true
}

func sudoku_solver(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := 49; k <= 57; k++ {
					if can_place(board, i, j, byte(k)) {
						board[i][j] = byte(k)

						if sudoku_solver(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	sudoku_solver(board)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}
