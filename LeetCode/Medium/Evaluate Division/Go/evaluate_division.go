package main

import "fmt"

//-----------------------------union find------------------------------------
type Pair struct {
	node string
	val  float64
}

func evaluate_dvision_union_find(equations [][]string, values []float64, queries [][]string) []float64 {
	parent := make(map[string]Pair)
	for i := 0; i < len(equations); i++ {
		union(equations[i][0], equations[i][1], values[i], &parent)
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		_, ok_a := parent[a]
		_, ok_b := parent[b]
		if !ok_a || !ok_b {
			res[i] = -1.0
		} else {
			p_a := find(a, &parent)
			p_b := find(b, &parent)
			if p_a.node != p_b.node {
				res[i] = -1.0
			} else {
				res[i] = p_a.val / p_b.val
			}
		}
	}
	return res
}

func union(a, b string, value float64, parent *map[string]Pair) {
	dividend := find(a, parent)
	divisor := find(b, parent)
	if dividend.node != divisor.node {
		(*parent)[dividend.node] = Pair{divisor.node, divisor.val * value / dividend.val}
	}
}

func find(s string, parent *map[string]Pair) Pair {
	if _, ok := (*parent)[s]; !ok {
		(*parent)[s] = Pair{s, 1.0}
	}
	p := (*parent)[s]
	if p.node != s {
		pair := find(p.node, parent)
		(*parent)[s] = Pair{pair.node, p.val * pair.val}
	}
	return (*parent)[s]

}

//-----------------------------union find------------------------------------

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
