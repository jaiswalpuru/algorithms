package main

import "fmt"

//--------------------dp----------------------
func bomb_enemy_dp(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	cnt, row_hits := 0, 0
	col_hits := make([]int, m)
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if c == 0 || grid[r][c-1] == 'W' {
				row_hits = 0
				for k := c; k < m; k++ {
					if grid[r][k] == 'W' {
						break
					}
					if grid[r][k] == 'E' {
						row_hits++
					}
				}
			}
			if r == 0 || grid[r-1][c] == 'W' {
				col_hits[c] = 0
				for k := r; k < n; k++ {
					if grid[k][c] == 'W' {
						break
					}
					if grid[k][c] == 'E' {
						col_hits[c]++
					}
				}
			}
			if grid[r][c] == '0' {
				cnt = max(cnt, row_hits+col_hits[c])
			}
		}
	}
	return cnt
}

//--------------------dp----------------------

//-------------brute force----------------------
func bomb_enemy(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				e := 0
				for k := j - 1; k >= 0; k-- {
					if grid[i][k] == 'W' {
						break
					}
					if grid[i][k] == 'E' {
						e++
					}
				}
				for k := j + 1; k < m; k++ {
					if grid[i][k] == 'W' {
						break
					}
					if grid[i][k] == 'E' {
						e++
					}
				}
				for k := i - 1; k >= 0; k-- {
					if grid[k][j] == 'W' {
						break
					}
					if grid[k][j] == 'E' {
						e++
					}
				}
				for k := i + 1; k < n; k++ {
					if grid[k][j] == 'W' {
						break
					}
					if grid[k][j] == 'E' {
						e++
					}
				}
				res = max(res, e)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------brute force------------------------

func main() {
	fmt.Println(bomb_enemy_dp([][]byte{
		{'0', 'E', '0', '0'}, {'E', '0', 'W', 'E'}, {'0', 'E', '0', '0'},
	}))
}
