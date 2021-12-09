/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that
no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.'
both indicate a queen and an empty space, respectively.

Example 1 :
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
*/

package main

import "fmt"

func is_safe(board [][]int, row, col int) bool {
	//check the left side of the column
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

	n := len(board)
	//check the lower diagonal
	for i, j := row, col; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func _n_queens(board *[][]int, col int) bool {
	//if all the queens are places return
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
		fmt.Println("Solution doesn't exist")
	} else {
		res := make([]string, 0)
		for i := 0; i < n; i++ {
			m := len(board[i])
			temp := ""
			for j := 0; j < m; j++ {
				if board[i][j] == 0 {
					temp += "."
				} else {
					temp += "Q"
				}
			}
			res = append(res, temp)
		}
		fmt.Println(res)
	}
}

func main() {
	n_queens(4)
	n_queens(1)
}
