package main

import "fmt"

func find_safe_states(graph [][]int) []int {
	n := len(graph)
	color := make([]int, n)
	res := []int{}

	for i := 0; i < n; i++ {
		if dfs(graph, color, i) {
			res = append(res, i)
		}
	}
	return res
}

//0-> white 1-> gray 2-> black
func dfs(graph [][]int, color []int, node int) bool {
	if color[node] > 0 {
		return color[node] == 2
	}

	color[node] = 1 //while entering we paint the node as gray and if in recursion gray color is encountered then cycle is present
	for i := 0; i < len(graph[node]); i++ {
		if color[graph[node][i]] == 2 {
			continue
		}
		if color[graph[node][i]] == 1 || !dfs(graph, color, graph[node][i]) {
			return false
		}
	}
	color[node] = 2
	return true
}

func main() {
	fmt.Println(find_safe_states([][]int{
		{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {},
	}))
}
