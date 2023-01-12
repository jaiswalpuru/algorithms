package main

import "fmt"

/*
	Tarjan algorith to return the number
	of scc present in graph, this algorithm
	can help in solving variety of problems,
	eg : articulation points, bridges etc.
	For a detailed explanation of this algo
	refer : William Fiset

	Here I will be using discovery time instead
	of id array for easy understanding.
*/

//return the low link values
func scc(edges [][]int, n int) {
	g := make_graph(edges)
	dis_time := make([]int, n) //discovery time, earliest time when the node was discovered
	low := make([]int, n)      //low time -> id of the node
	sccs := make([]int, n)
	for i := 0; i < n; i++ {
		dis_time[i] = -1 //marking as not visited
	}
	stack := []int{} //this will be used to keep track of the nodes which form a component
	time := 0
	scc_cnt := 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if dis_time[i] == -1 {
			dfs(&dis_time, &low, &sccs, &stack, &visited, &time, i, g, &scc_cnt)
		}
	}
	fmt.Println("Total number of scc : ", scc_cnt)
	mp := make(map[int][]int)
	for i := 0; i < len(sccs); i++ {
		mp[sccs[i]] = append(mp[sccs[i]], i)
	}
	fmt.Println(mp)
}

func dfs(dis_time, low, sccs, stack *[]int, visited *[]bool, time *int, u int, g map[int][]int, scc_cnt *int) {
	*time++
	(*dis_time)[u] = *time
	(*low)[u] = *time
	(*visited)[u] = true
	*stack = append(*stack, u) //the root is pushed to stack
	//traverse the nodes adjacent to u
	for i := 0; i < len(g[u]); i++ {
		v := g[u][i]
		if (*dis_time)[v] == -1 {
			dfs(dis_time, low, sccs, stack, visited, time, v, g, scc_cnt)
		}
		//callback from dfs
		if (*visited)[v] {
			(*low)[u] = min((*low)[u], (*low)[v])
		}
	}
	//if at the root, if the discovery time and low link value are same we know that a scc is found
	if (*dis_time)[u] == (*low)[u] {
		for len(*stack) > 0 {
			node := (*stack)[len(*stack)-1]
			*stack = (*stack)[:len(*stack)-1]
			(*visited)[node] = false //mark it as false so that it doesn't mess with other components
			(*sccs)[node] = *scc_cnt
			if node == u { //pop until root has been found
				break
			}
		}
		*scc_cnt++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
	}
	return g
}

func main() {
	g1 := [][]int{
		{6, 0}, {6, 2}, {3, 4}, {6, 4},
		{2, 0}, {0, 1}, {4, 5}, {5, 6},
		{3, 7}, {7, 5}, {1, 2}, {7, 3},
		{5, 0},
	}
	scc(g1, 8)
}
