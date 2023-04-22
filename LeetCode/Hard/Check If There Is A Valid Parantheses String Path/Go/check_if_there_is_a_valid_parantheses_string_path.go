package main

import "fmt"

func is_safe(i, j, n, m int) bool {
	return (i >= 0 && j >= 0 && i < n && j < m)
}

//------------------brute force trying all possible paths------------------------------------
func _check(i, j, n, m int, grid [][]byte, open, close int) bool {
	if is_safe(i, j, n, m) {
		if close > open {
			return false
		}
		if grid[i][j] == '(' {
			open++
		} else {
			close++
		}
		if i == n-1 && j == m-1 && close == open {
			return true
		}
		return _check(i+1, j, n, m, grid, open, close) || _check(i, j+1, n, m, grid, open, close)
	}
	return false
}

func check_if_there_is_valid_parantheses_string_path(grid [][]byte) bool {
	n, m := len(grid), len(grid[0])
	return _check(0, 0, n, m, grid, 0, 0)
}

//------------------------------------------------------------------------------------------

//-----------------------------------memoization--------------------------------------
func check_if_there_is_valid_parantheses_string_path_memo(grid [][]byte) bool {
	n, m := len(grid), len(grid[0])
	p := n + m - 2
	memo := make([][][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = make([]int, p+1)
			for k := 0; k <= p; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	return _check_memo(grid, 0, 0, n, m, &memo, 0)
}

func _check_memo(grid [][]byte, i, j, n, m int, memo *[][][]int, open int) bool {
	if !is_safe(i, j, n, m) {
		return false
	}

	if open < 0 || open > (n+m-2) {
		return false
	}

	if (*memo)[i][j][open] != -1 {
		return (*memo)[i][j][open] == 1
	}

	if i == n-1 && j == m-1 {
		if grid[i][j] == '(' {
			open++
		} else {
			open--
		}
		return open == 0
	}

	var v bool
	if grid[i][j] == '(' {
		v = _check_memo(grid, i+1, j, n, m, memo, open+1) || _check_memo(grid, i, j+1, n, m, memo, open+1)
	} else {
		v = _check_memo(grid, i+1, j, n, m, memo, open-1) || _check_memo(grid, i, j+1, n, m, memo, open-1)
	}
	if v {
		(*memo)[i][j][open] = 1
	} else {
		(*memo)[i][j][open] = 0
	}
	return v
}

//------------------------------------------------------------------------------------

func main() {
	fmt.Println(check_if_there_is_valid_parantheses_string_path_memo([][]byte{
		{'(', '(', '('}, {')', '(', ')'}, {'(', '(', ')'}, {'(', '(', ')'},
	}))
}
