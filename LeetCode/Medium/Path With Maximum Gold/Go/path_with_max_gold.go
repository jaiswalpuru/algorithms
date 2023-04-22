package main

import (
	"fmt"
)

func is_safe_to_walk(i, j, n, m int, visited [][]bool, arr [][]int) bool {
	if i < 0 || j < 0 || j >= m || i >= n || visited[i][j] || arr[i][j] == 0 {
		return false
	}
	return true
}

var direction = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func get_maximum_gold(arr [][]int) int {
	n, m := len(arr), len(arr[0])

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	max_gold := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] != 0 {
				max_gold = max(max_gold, _get_max_gold(i, j, n, m, arr, &visited))
			}
		}
	}
	return max_gold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _get_max_gold(i, j, n, m int, arr [][]int, visited *[][]bool) int {
	if !is_safe_to_walk(i, j, n, m, *visited, arr) {
		return 0
	}
	sum := 0
	(*visited)[i][j] = true
	for k := 0; k < 4; k++ {
		curr_i, curr_j := i+direction[k][0], j+direction[k][1]
		sum = max(sum, arr[i][j]+_get_max_gold(curr_i, curr_j, n, m, arr, visited))
	}
	(*visited)[i][j] = false
	return sum
}

func main() {
	fmt.Println(get_maximum_gold([][]int{
		{0, 6, 0}, {5, 8, 7}, {0, 9, 0},
	}))
}
