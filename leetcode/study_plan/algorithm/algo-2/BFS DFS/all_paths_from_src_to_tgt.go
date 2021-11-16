/*
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0
to node n - 1 and return them in any order.

The graph is given as follows: graph[i] is a list of all nodes you can visit from node i
(i.e., there is a directed edge from node i to node graph[i][j]).
*/

package main

import "fmt"

func all_path(graph [][]int) [][]int {
	g := make(map[int][]int)
	n := len(graph)
	for i := 0; i < n; i++ {
		g[i] = append(g[i], graph[i]...)
	}

	result := [][]int{}
	q := make([][]int, 0)
	q = append(q, []int{0})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		lastVertex := curr[len(curr)-1]
		for i := 0; i < len(g[lastVertex]); i++ {
			temp := curr
			temp = append(temp, g[lastVertex][i])
			if g[lastVertex][i] == n-1 {
				result = append(result, append([]int{}, temp...))
			} else {
				q = append(q, append([]int{}, temp...))
			}
		}
	}
	return result
}

func main() {
	// fmt.Println(all_path([][]int{
	// 	{1, 2}, {3}, {3}, {},
	// }))
	// fmt.Println(all_path([][]int{
	// 	{4, 3, 1}, {3, 2, 4}, {3}, {4}, {},
	// }))
	// fmt.Println(all_path([][]int{
	// 	{1}, {},
	// }))
	// fmt.Println(all_path([][]int{
	// 	{1, 2, 3}, {2}, {3}, {},
	// }))
	// fmt.Println(all_path([][]int{
	// 	{1, 3}, {2}, {3}, {},
	// }))
	fmt.Println(all_path([][]int{
		{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {},
	}))
}
