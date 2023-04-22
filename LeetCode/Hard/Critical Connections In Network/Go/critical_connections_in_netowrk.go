package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	g := make(map[int][]int)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func dfs(node int, depth int, rank *map[int]int, cycle *map[[2]int]struct{}, graph map[int][]int, n int) int {
	if (*rank)[node] >= 0 {
		return (*rank)[node]
	}

	(*rank)[node] = depth
	min_depth := n

	for i := 0; i < len(graph[node]); i++ {
		curr := graph[node][i]
		if (*rank)[curr] == depth-1 {
			continue
		}
		min_depth_temp := dfs(curr, depth+1, rank, cycle, graph, n)
		if min_depth_temp <= depth {
			(*cycle)[[2]int{node, curr}] = struct{}{}
		}
		min_depth = min(min_depth, min_depth_temp)
	}
	return min_depth
}

func critical_connections(n int, connections [][]int) [][]int {
	g := make_graph(connections)

	ranks := make(map[int]int)
	cycle := make(map[[2]int]struct{})

	for i := 0; i < n; i++ {
		ranks[i] = -2
	}

	dfs(0, 0, &ranks, &cycle, g, n)
	res := make([][]int, 0)

	for _, conn := range connections {
		if _, ok := cycle[[2]int{conn[0], conn[1]}]; ok {
			continue
		}
		if _, ok := cycle[[2]int{conn[1], conn[0]}]; ok {
			continue
		}
		res = append(res, conn)
	}
	return res
}

//---------------------------------using tarjan's algorithm----------------------------------
type Pair struct { discovery_time, low_number int }

func critical_connections_tarjan(n int, conn [][]int) [][]int {
	g := make_graph(conn)

	nodes := make([]Pair, n)
	res :=[][]int{}
	discovery_time := 0
	tarjan(g, nodes, &discovery_time, &res, 0, -1)
	return res
}

func tarjan(graph map[int][]int, nodes []Pair, discovery_time *int, res *[][]int, curr, parent int) {
	*discovery_time++
	nodes[curr].discovery_time = *discovery_time
	nodes[curr].low_number = *discovery_time

	for i:=0;i<len(graph[curr]);i++{
		v := graph[curr][i]
		if v == parent {
			continue
		}
		if nodes[v].discovery_time == 0 {
			tarjan(graph, nodes, discovery_time, res, v, curr)
			nodes[curr].low_number = min(nodes[curr].low_number, nodes[v].low_number)
			if nodes[v].low_number > nodes[curr].discovery_time {
				*res = append(*res, []int{curr, v})
			}
		}
		nodes[curr].low_number = min(nodes[curr].low_number, nodes[v].discovery_time)
	}
}
//---------------------------------using tarjan's algorithm----------------------------------

func main() {
	fmt.Println(critical_connections_tarjan(4, [][]int{
		{0, 1}, {1, 2}, {2, 0}, {1, 3},
	}))
}
