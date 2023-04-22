package main

import "fmt"

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

//----------------------dfs--------------------------------------
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

//----------------------dfs--------------------------------------

//----------------------bfs--------------------------------------
func pacific_atlantic_water_flow_bfs(height [][]int) [][]int {
	if len(height) == 0 {
		return nil
	}
	n, m := len(height), len(height[0])
	pac_q, atl_q := [][]int{}, [][]int{}
	for i := 0; i < n; i++ {
		pac_q = append(pac_q, []int{i, 0})
		atl_q = append(atl_q, []int{i, n - 1})
	}
	for j := 0; j < m; j++ {
		pac_q = append(pac_q, []int{0, j})
		atl_q = append(atl_q, []int{n - 1, j})
	}
	pac_visited := make([][]bool, n)
	atl_visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		pac_visited[i] = make([]bool, m)
		atl_visited[i] = make([]bool, m)
	}
	bfs(&pac_visited, height, pac_q)
	bfs(&atl_visited, height, atl_q)
	res := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pac_visited[i][j] && atl_visited[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func bfs(visited *[][]bool, height [][]int, q [][]int) {
	n, m := len(height), len(height[0])
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		(*visited)[curr[0]][curr[1]] = true
		for i := 0; i < 4; i++ {
			x, y := curr[0]+dir[i][0], curr[1]+dir[i][1]
			if x < 0 || y < 0 || x >= n || y >= m {
				continue
			}
			if (*visited)[x][y] {
				continue
			}
			if height[x][y] < height[curr[0]][curr[1]] {
				continue
			}
			q = append(q, []int{x, y})
		}
	}
}

//----------------------bfs--------------------------------------

func main() {
	fmt.Println(pacific_atlantic_water_flow_bfs([][]int{
		{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4},
	}))
}
