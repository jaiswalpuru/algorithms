/*
	You have an N by N board. Write a function that, given N,
	returns the number of possible arrangements of the board where N queens can be placed on the board
	without threatening each other, i.e. no two queens share the same row, column, or diagonal.
*/

package main

import (
	"fmt"
	"math"
)

//board is 1D array which represents the col position of the queen

func n_queens(n int, board []int) int {
	if n == len(board) {
		return 1
	}

	count := 0
	for i := 0; i < n; i++ {
		board = append(board, i)

		if is_valid(board) {
			count += n_queens(n, board)
		}
		board = board[:len(board)-1]
	}
	return count
}

func is_valid(board []int) bool {
	queen_row, queen_col := len(board)-1, board[len(board)-1]

	for i := 0; i < len(board)-1; i++ {
		diff := int(math.Abs(float64(queen_col - board[i])))
		if diff == 0 || diff == queen_row-i {
			return false
		}
	}
	return true
}

func main() {
	n := 4
	arr := []int{}
	fmt.Println(n_queens(n, arr))
}
