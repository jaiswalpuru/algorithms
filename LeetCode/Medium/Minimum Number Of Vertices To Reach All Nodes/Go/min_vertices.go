package main

import "fmt"

func min_vertices(n int, arr [][]int) []int {
	indegree := make(map[int]int, n)

	m := len(arr)
	for i := 0; i < m; i++ {
		indegree[arr[i][1]]++
	}

	res := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	fmt.Println(min_vertices(6, [][]int{
		{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2},
	}))
}
