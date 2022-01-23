package main

import "fmt"

func make_graph(edges [][]string, values []float64) (map[string]map[string]float64, map[string]bool) {
	graph := make(map[string]map[string]float64)
	visited := make(map[string]bool)
	n := len(edges)

	for i := 0; i < n; i++ {
		if _, ok := graph[edges[i][0]]; !ok {
			graph[edges[i][0]] = make(map[string]float64)
		}
		graph[edges[i][0]][edges[i][1]] = values[i]
		if _, ok := graph[edges[i][1]]; !ok {
			graph[edges[i][1]] = make(map[string]float64)
		}
		graph[edges[i][1]][edges[i][0]] = 1 / values[i]

		visited[edges[i][0]] = false
		visited[edges[i][1]] = false
	}

	return graph, visited
}

func dfs(graph map[string]map[string]float64, visited map[string]bool, dividend, divisor string) float64 {
	visited[dividend] = true
	l := graph[dividend]

	if _, ok := l[divisor]; ok {
		return graph[dividend][divisor]
	}

	for k, v := range l {
		if visited[k] {
			continue
		}

		res := dfs(graph, visited, k, divisor)
		if res != -1 {
			return v * res
		}
	}
	return -1
}

func evaluate_dvision(equations [][]string, values []float64, queries [][]string) []float64 {
	graph, visited := make_graph(equations, values)
	n := len(queries)
	res := make([]float64, n)

	for i := 0; i < n; i++ {
		dividend := queries[i][0]
		divisor := queries[i][1]

		_, ok_1 := graph[dividend]
		_, ok_2 := graph[divisor]

		if !ok_1 || !ok_2 {
			res[i] = -1
			continue
		}

		if dividend == divisor {
			res[i] = 1
			continue
		}

		for k := range visited {
			visited[k] = false
		}

		res[i] = dfs(graph, visited, dividend, divisor)
	}

	return res
}

func main() {
	fmt.Println(evaluate_dvision([][]string{
		{"a", "b"}, {"b", "c"},
	}, []float64{2.0, 3.0}, [][]string{
		{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"},
	}))
}
