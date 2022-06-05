package main

import "fmt"

func n_queens_two(n int) int {
	if n == 1 {
		return 1
	}
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
	}
	res := [][]string{}
	_backtrack(&b, 0, n, &res)
	return len(res)
}

func _backtrack(b *[][]int, col, n int, ans *[][]string) {
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
		*ans = append(*ans, temp)
		return
	}

	for row := 0; row < n; row++ {
		if is_safe(row, col, *b, n) {
			(*b)[row][col] = 1
			_backtrack(b, col+1, n, ans)
			(*b)[row][col] = 0
		}
	}
}

func is_safe(row, col int, b [][]int, n int) bool {
	r, c := row, col
	//check upped diagonal
	for r >= 0 && c >= 0 {
		if b[r][c] == 1 {
			return false
		}
		r--
		c--
	}

	//check left col
	r, c = row, col
	for c >= 0 {
		if b[r][c] == 1 {
			return false
		}
		c--
	}

	//check left diagonal
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
	fmt.Println(n_queens_two(4))
}
