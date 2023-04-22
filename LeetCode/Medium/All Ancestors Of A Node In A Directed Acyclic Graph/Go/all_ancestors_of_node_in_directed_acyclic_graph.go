package main

import "fmt"

func make_graph(edges [][]int) map[int][]int{
	graph := make(map[int][]int)
	n:=len(edges)
	for i:=0;i<n;i++{
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
	}
	return graph
}

func all_ancestors_of_node_in_dag(n int, edges [][]int) [][]int {
	graph := make_graph(edges)

	res := [][]int{}
	g := make(map[int][]int)
	for i:=0;i<n;i++{
		q := []int{i}
		visited := make(map[int]bool)
		for len(q) > 0{
			curr := q[0]
			q= q[1:]
			for j:=0;j<len(graph[curr]);j++{
				if !visited[graph[curr][j]]{
					g[graph[curr][j]] = append(g[graph[curr][j]], i)
					visited[graph[curr][j]] = true
					q =append(q, graph[curr][j])
				}
			}
		}
	}
	for i:=0;i<n;i++{
		res = append(res, g[i])
	}
	return res
}

func main () {
	fmt.Println(all_ancestors_of_node_in_dag(8, [][]int{
		{0,3},{0,4},{1,3},{2,4},{2,7},{3,5},{3,6},{3,7},{4,6},
	}))
}