package main

import "fmt"

func count_server_that_communicate(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	connected_servers := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				q := [][]int{{i, j}}
				cnt := 0
				for len(q) > 0 {
					curr_x, curr_y := q[0][0], q[0][1]
					q = q[1:]
					for k := curr_x; k < n; k++ {
						if grid[k][curr_y] == 1 && !visited[k][curr_y] {
							cnt += 1
							visited[k][curr_y] = true
							q = append(q, []int{k, curr_y})
						}
					}
					for k := curr_y; k < m; k++ {
						if grid[curr_x][k] == 1 && !visited[curr_x][k] {
							cnt += 1
							visited[curr_x][k] = true
							q = append(q, []int{curr_x, k})
						}
					}
					for k := 0; k < curr_x; k++ {
						if grid[k][curr_y] == 1 && !visited[k][curr_y] {
							cnt += 1
							visited[k][curr_y] = true
							q = append(q, []int{k, curr_y})
						}
					}
					for k := 0; k < curr_y; k++ {
						if grid[curr_x][k] == 1 && !visited[curr_x][k] {
							cnt += 1
							visited[curr_x][k] = true
							q = append(q, []int{curr_x, k})
						}
					}
				}
				if cnt > 1 {
					connected_servers += cnt
				}
			}
		}
	}
	return connected_servers
}

func main() {
	fmt.Println(count_server_that_communicate([][]int{
		{0, 0, 1, 0, 1}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}, {1, 0, 0, 1, 1}, {0, 0, 1, 1, 0},
	}))
}
