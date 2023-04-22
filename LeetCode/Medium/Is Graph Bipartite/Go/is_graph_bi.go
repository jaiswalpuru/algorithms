package main

import "fmt"

func is_bi(edges [][]int) bool {

	n := len(edges)

	// 0-> not traversed, 1-> belongs to set 1, 2-> belongs to set 2
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			q := []int{i}
			for len(q) > 0 {
				curr := q[0]
				q = q[1:]
				for j := 0; j < len(edges[curr]); j++ {
					if visited[curr] == visited[edges[curr][j]] {
						return false
					}
					if visited[edges[curr][j]] == 0 {
						q = append(q, edges[curr][j])
						if visited[curr] == 1 {
							visited[edges[curr][j]] = 2
						} else {
							visited[edges[curr][j]] = 1
						}
					}
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(is_bi([][]int{
		{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2},
	}))
}
