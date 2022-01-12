/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a
queen and an empty space, respectively.
*/
package main

import "fmt"

func n_queens(n int) [][]string {
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
	}
	ans := [][]string{}
	_n_queens(&b, 0, n, &ans)
	return ans
}

func _n_queens(board *[][]int, col, n int, ans *[][]string) {
	if col >= n {
		temp := []string{}
		for i := 0; i < n; i++ {
			str := ""
			for j := 0; j < n; j++ {
				if (*board)[i][j] == 0 {
					str += "."
				} else {
					str += "Q"
				}
			}
			temp = append(temp, str)
		}
		*ans = append(*ans, temp)
		return
	}

	for row := 0; row < n; row++ {
		if is_safe(row, col, *board, n) {
			(*board)[row][col] = 1
			_n_queens(board, col+1, n, ans)
			(*board)[row][col] = 0
		}
	}
}

func is_safe(row, col int, board [][]int, n int) bool {
	r, c := row, col

	//check upper diagonal
	for r >= 0 && c >= 0 {
		if board[r][c] == 1 {
			return false
		}
		r--
		c--
	}

	//check all the columns before the current one
	r = row
	c = col
	for c >= 0 {
		if board[r][c] == 1 {
			return false
		}
		c--
	}

	//check all the left lower diagonal before the current one
	r = row
	c = col
	for r < n && c >= 0 {
		if board[r][c] == 1 {
			return false
		}
		r++
		c--
	}

	return true
}

func main() {
	fmt.Println(n_queens(4))
}
