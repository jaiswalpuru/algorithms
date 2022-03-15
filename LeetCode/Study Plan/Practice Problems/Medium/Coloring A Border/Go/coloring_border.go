package main

import "fmt"

func coloring_border(arr [][]int, r, c int, color int) [][]int {
	n, m := len(arr), len(arr[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	curr_val := arr[r][c]

	dfs(&arr, r, c, n, m, color, &visited, curr_val)
	return arr
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(arr *[][]int, i, j, n, m, color int, visited *[][]bool, curr int) {

	(*visited)[i][j] = true
	if i == 0 || j == 0 || i == n-1 || j == m-1 {
		(*arr)[i][j] = color
	} else {
		for k := 0; k < 4; k++ {
			x, y := i+directions[k][0], j+directions[k][1]
			if x >= 0 && y >= 0 && x < n && y < m && !(*visited)[x][y] && (*arr)[x][y] != curr {
				(*arr)[i][j] = color
			}
		}
	}

	for k := 0; k < 4; k++ {
		x, y := i+directions[k][0], j+directions[k][1]
		if x >= 0 && y >= 0 && x < n && y < m && !(*visited)[x][y] && (*arr)[x][y] == curr {
			dfs(arr, x, y, n, m, color, visited, curr)
		}
	}

}

func main() {
	fmt.Println(coloring_border([][]int{
		{1, 2, 1, 2, 1, 2}, {2, 2, 2, 2, 1, 2}, {1, 2, 2, 2, 1, 2},
	}, 1, 3, 1))
}
