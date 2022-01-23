package main

import "fmt"

func num_provinces(is_connected [][]int) int {
	n := len(is_connected)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if is_connected[i][j] == 1 && visited[i][j] != true {
				cnt++
				visited[i][j] = true
				dfs(i, j, is_connected, &visited)
			}
		}
	}
	return cnt
}

func dfs(i, j int, is_connected [][]int, visited *[][]bool) {
	(*visited)[i][j] = true
	for k := 0; k < len(is_connected); k++ {
		if !(*visited)[j][k] && is_connected[j][k] == 1 {
			dfs(j, k, is_connected, visited)
		}
	}
}

func main() {
	fmt.Println(num_provinces([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))
}
