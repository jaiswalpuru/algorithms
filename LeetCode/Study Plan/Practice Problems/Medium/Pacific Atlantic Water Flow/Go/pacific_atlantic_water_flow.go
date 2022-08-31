package main

import "fmt"

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func pacific_atlantic_water_flow(height [][]int) [][]int {
	if len(height) == 0 {
		return nil
	}
	res := [][]int{}
	n, m := len(height), len(height[0])
	pacific := make([][]bool, n)
	atlantic := make([][]bool, n)
	for i := 0; i < n; i++ {
		pacific[i] = make([]bool, m)
		atlantic[i] = make([]bool, m)
	}
	///loop through each cell adjacent to oceans
	for i := 0; i < n; i++ {
		dfs(i, 0, &pacific, height)
		dfs(i, m-1, &atlantic, height)
	}
	for i := 0; i < m; i++ {
		dfs(0, i, &pacific, height)
		dfs(n-1, i, &atlantic, height)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func dfs(row, col int, visited *[][]bool, height [][]int) {
	(*visited)[row][col] = true
	for i := 0; i < 4; i++ {
		r, c := row+dir[i][0], col+dir[i][1]
		if r < 0 || c < 0 || r >= len(height) || c >= len(height[0]) || (*visited)[r][c] {
			continue
		}
		if height[r][c] < height[row][col] {
			continue
		}
		dfs(r, c, visited, height)
	}
}

func main() {
	fmt.Println(pacific_atlantic_water_flow([][]int{
		{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4},
	}))
}
