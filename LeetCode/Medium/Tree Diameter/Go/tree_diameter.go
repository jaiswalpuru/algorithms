package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	return graph
}

func tree_diameter(edges [][]int) int {
	graph := make_graph(edges)
	visited := make(map[int]bool)

	a, b, d_a, d_b := 0, 0, 0, 0

	dfs_get_deepest_node(graph, 0, &a, &d_a, 0, &visited) //update first deep node
	visited = make(map[int]bool)
	dfs_get_deepest_node(graph, a, &b, &d_b, 0, &visited) // update the second deep node
	dis := 0
	type Pair struct {
		node int
		dis  int
	}

	q := []Pair{{a, 0}}
	visited = make(map[int]bool)
	for len(q) > 0 {
		curr := q[0].node
		dis = q[0].dis
		q = q[1:]
		if curr == b {
			break
		}
		visited[curr] = true
		for i := 0; i < len(graph[curr]); i++ {
			if !visited[graph[curr][i]] {
				q = append(q, Pair{graph[curr][i], dis + 1})
			}
		}
	}

	return dis
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func dfs_get_deepest_node(graph map[int][]int, start int, deep, d_a *int, temp int, visited *map[int]bool) {
	(*visited)[start] = true
	for i := 0; i < len(graph[start]); i++ {
		if !(*visited)[graph[start][i]] {
			if temp >= *d_a {
				*d_a = temp
				*deep = graph[start][i]
			}
			dfs_get_deepest_node(graph, graph[start][i], deep, d_a, temp+1, visited)
		}
	}
}

func main() {
	fmt.Println(tree_diameter([][]int{
		{0, 1}, {0, 2},
	}))
}
