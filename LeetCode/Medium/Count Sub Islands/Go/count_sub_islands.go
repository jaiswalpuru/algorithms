package main

import "fmt"

func count_sub_islands(grid_one [][]int, grid_two [][]int) int {
	n, m := len(grid_one), len(grid_one[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	cnt := 0
	inner_cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && grid_two[i][j] == 1 {
				inner_cnt = 1
				dfs(grid_one, grid_two, &visited, i, j, n, m, &inner_cnt)
				cnt += inner_cnt
			}
		}
	}
	return cnt
}

func dfs(grid_one, grid_two [][]int, visited *[][]bool, i, j, n, m int, inner_cnt *int) {
	if i >= n || j >= m || i < 0 || j < 0 || (*visited)[i][j] || grid_two[i][j] == 0 {
		return
	}
	if grid_one[i][j] == 0 {
		*inner_cnt = 0
		return
	}
	(*visited)[i][j] = true
	dfs(grid_one, grid_two, visited, i+1, j, n, m, inner_cnt)
	dfs(grid_one, grid_two, visited, i-1, j, n, m, inner_cnt)
	dfs(grid_one, grid_two, visited, i, j+1, n, m, inner_cnt)
	dfs(grid_one, grid_two, visited, i, j-1, n, m, inner_cnt)
}

func main() {
	fmt.Println(count_sub_islands([][]int{
		{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1},
	}, [][]int{
		{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0},
	}))
}
