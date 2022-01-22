package main

import (
	"fmt"
	"math"
)

//-------------------------------------------dirty brute force-------------------------------------------
type Pair struct {
	val int
	ht  int
}

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func min_ht_trees(n int, edges [][]int) []int {
	g := make_graph(edges)
	min_ht_node := []int{}
	min_ht := math.MaxInt64
	for i := 0; i < n; i++ {
		ht := get_ht(i, g)
		if min_ht > ht {
			min_ht = ht
			min_ht_node = []int{i}
		} else if min_ht == ht {
			min_ht_node = append(min_ht_node, i)
		}
	}
	return min_ht_node
}

func get_ht(start_node int, graph map[int][]int) int {
	q := []Pair{{val: start_node, ht: 1}}
	visited := make(map[int]int)
	visited[start_node] = -1
	var ht int
	for len(q) > 0 {
		curr := q[0].val
		ht = q[0].ht
		q = q[1:]
		for i := 0; i < len(graph[curr]); i++ {
			if visited[curr] == graph[curr][i] {
				continue
			}
			visited[graph[curr][i]] = curr
			q = append(q, Pair{val: graph[curr][i], ht: ht + 1})
		}
	}
	return ht
}

//--------------------------------------------------------------------------------------
func get_min_ht(n int, edges [][]int) []int {
	graph := make([][]int, n)
	degree := make([]int, n)
	m := len(edges)
	for i := 0; i < m; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
		degree[edges[i][0]]++
		degree[edges[i][1]]++
	}

	node_cnt := n
	for node_cnt > 2 {
		q := []int{}

		for i := 0; i < n; i++ {
			if degree[i] == 1 {
				q = append(q, i)
				degree[i] = -1
				node_cnt--
			}
		}

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			for i := 0; i < len(graph[curr]); i++ {
				degree[graph[curr][i]]--
			}
		}
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if degree[i] == 1 || degree[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(get_min_ht(6, [][]int{
		{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4},
	}))
	fmt.Println(get_min_ht(3, [][]int{{0, 1}, {0, 2}}))
}
