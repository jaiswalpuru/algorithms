package main

import "fmt"

func max_number_of_accepted_invitation(grid [][]int) int {
	match := make([]int, 201)
	n, m := len(grid), len(grid[0])

	cnt := 0
	for i := 1; i <= n; i++ {
		visited := make([]int, m+1)
		if max_bipartite(i, m, grid, &visited, &match) {
			cnt++
		}
	}
	return cnt
}

func max_bipartite(boy, m int, grid [][]int, visited, match *[]int) bool {
	for girl := 1; girl <= m; girl++ {
		// if the boy can invite the girl and the girl hasn't been invited yet.
		if grid[boy-1][girl-1] == 1 && (*visited)[girl] == 0 {
			(*visited)[girl] = 1
			if (*match)[girl] == 0 || max_bipartite((*match)[girl], m, grid, visited, match) {
				(*match)[girl] = boy
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(max_number_of_accepted_invitation([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{0, 0, 1},
	}))
}
