package main

import "fmt"

//-------------------------------------------------------------------------(refer LC)
func shift_2d_grid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	for k > 0 {

		n_grid := make([][]int, n)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				n_grid[i] = make([]int, m)
			}
		}
		//move everything which is not in the last column
		for i := 0; i < n; i++ {
			for j := 0; j < m-1; j++ {
				n_grid[i][j+1] = grid[i][j]
			}
		}

		//move everything in from last column but the last row
		for i := 0; i < n-1; i++ {
			n_grid[i+1][0] = grid[i][m-1]
		}

		n_grid[0][0] = grid[n-1][m-1]
		grid = n_grid
		k--
	}
	return grid
}

//--------------------------------------------------------------------------------

//--------------------------------method 2----------------------------------------
func shift_2d_grid_m2(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	for k > 0 {
		prev := grid[n-1][m-1]
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				temp := grid[i][j]
				grid[i][j] = prev
				prev = temp
			}
		}
		k--
	}
	return grid
}

//--------------------------------method 2----------------------------------------

func main() {
	fmt.Println(shift_2d_grid_m2([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}, 1))
}
