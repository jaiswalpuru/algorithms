package main

import (
	"fmt"
	"strconv"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

//-------------------------this solution is accepted but not space optimized------------------------------

func are_equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	n := len(a)
	x_diff := abs(a[0][0] - b[0][0])
	y_diff := abs(a[0][1] - b[0][1])
	for i := 0; i < n; i++ {
		if x_diff != abs(a[i][0]-b[i][0]) || y_diff != abs(a[i][1]-b[i][1]) {
			return false
		}
	}
	return true
}

func number_of_distinct_islands(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	cnt := 0
	points := [][][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				temp := [][]int{}
				dfs_not_ff(grid, &visited, i, j, n, m, &temp)
				cnt++
				if len(temp) != 0 {
					points = append(points, temp)
				}
			}
		}
	}

	v := make(map[string]bool)
	for i := 0; i < len(points); i++ {
		v[ptos(points[i])] = true
		for j := i + 1; j < len(points); j++ {
			if v[ptos(points[j])] {
				continue
			}
			if are_equal(points[i], points[j]) {
				v[ptos(points[j])] = true
				cnt--
			}
		}
		if cnt <= 0 {
			cnt = 1
		}
	}
	return cnt
}

func ptos(p [][]int) string {
	n := len(p)
	str := ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa(p[i][0]) + ":" + strconv.Itoa(p[i][1])
	}
	return str
}

func dfs_not_ff(grid [][]int, visited *[][]bool, i, j, n, m int, temp *[][]int) {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 0 || (*visited)[i][j] {
		return
	}

	(*visited)[i][j] = true
	*temp = append(*temp, []int{i, j})
	dfs_not_ff(grid, visited, i+1, j, n, m, temp)
	dfs_not_ff(grid, visited, i-1, j, n, m, temp)
	dfs_not_ff(grid, visited, i, j+1, n, m, temp)
	dfs_not_ff(grid, visited, i, j-1, n, m, temp)
}

//--------------------------------------------------------------------------------------------------

//-------------------------------------a better solution--------------------------------------------
func number_of_distinct_islands_eff(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	set := make(map[string]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp := [][]int{}
			curr_row, curr_col := i, j
			dfs(grid, i, j, curr_row, curr_col, n, m, &visited, &temp)
			if len(temp) != 0 {
				set[ptos(temp)] = struct{}{}
			}
		}
	}
	return len(set)
}

func dfs(grid [][]int, i, j, curr_row, curr_col, n, m int, visited *[][]bool, temp *[][]int) {
	if i < 0 || j < 0 || i >= n || j >= m || (*visited)[i][j] || grid[i][j] == 0 {
		return
	}
	(*visited)[i][j] = true
	*temp = append(*temp, []int{i - curr_row, j - curr_col})
	dfs(grid, i+1, j, curr_row, curr_col, n, m, visited, temp)
	dfs(grid, i-1, j, curr_row, curr_col, n, m, visited, temp)
	dfs(grid, i, j+1, curr_row, curr_col, n, m, visited, temp)
	dfs(grid, i, j-1, curr_row, curr_col, n, m, visited, temp)
}

//-------------------------------------a better solution--------------------------------------------

func main() {
	fmt.Println(number_of_distinct_islands_eff([][]int{
		{1, 1, 0, 1, 1}, {1, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {1, 1, 0, 1, 1},
	}))
}
