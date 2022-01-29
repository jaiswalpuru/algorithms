package main

import "fmt"

func town_judge(n int, graph [][]int) int {
	//just calculate the outdegree and indegrees

	m := len(graph)
	indegree := make(map[int]int)
	outdegree := make(map[int]int)
	for i := 0; i < m; i++ {
		outdegree[graph[i][0]]++
		indegree[graph[i][1]]++
	}
	for i := 1; i <= n; i++ {
		if indegree[i] == n-1 && outdegree[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(town_judge(2, [][]int{{1, 2}}))
}
