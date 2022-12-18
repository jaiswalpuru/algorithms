package main

import "fmt"

func n_queens(n int) [][]string {
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
	}
	res := [][]string{}
	backtrack(&b, &res, 0, n)
	return res
}

func backtrack(b *[][]int, res *[][]string, col int, n int) {
	if col >= n {
		temp := []string{}
		for i := 0; i < n; i++ {
			str := ""
			for j := 0; j < n; j++ {
				if (*b)[i][j] == 1 {
					str += "Q"
				} else {
					str += "."
				}
			}
			temp = append(temp, str)
		}
		*res = append(*res, temp)
		return
	}
	for r := 0; r < n; r++ {
		if is_safe(r, col, *b, n) {
			(*b)[r][col] = 1
			backtrack(b, res, col+1, n)
			(*b)[r][col] = 0
		}
	}
}

func is_safe(row, col int, b [][]int, n int) bool {
	r, c := row, col
	for r >= 0 && c >= 0 {
		if b[r][c] == 1 {
			return false
		}
		r--
		c--
	}
	r, c = row, col
	for c >= 0 {
		if b[r][c] == 1 {
			return false
		}
		c--
	}
	r, c = row, col
	for r < n && c >= 0 {
		if b[r][c] == 1 {
			return false
		}
		r++
		c--
	}
	return true
}

func main() {
	fmt.Println(n_queens(8))
}
