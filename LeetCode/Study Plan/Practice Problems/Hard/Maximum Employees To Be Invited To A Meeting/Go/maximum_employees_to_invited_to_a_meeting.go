package main

import "fmt"

func maximum_employees_to_invited_to_a_meeting(favorite []int) int {
	g := make_graph(favorite)
	n := len(favorite)
	q := []int{}
	visited := make([]bool, n)
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		indegree[i] = len(g[i])
	}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
			visited[i] = true
		}
	}
	path := make([]int, n) //path[i], longest path leading to i
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		curr_fav := favorite[curr]
		path[curr_fav] = max(path[curr_fav], 1+path[curr])
		indegree[curr_fav]--
		if indegree[curr_fav] == 0 {
			visited[curr_fav] = true
			q = append(q, curr_fav)
		}
	}
	//the not visited nodes are cycles
	res, res_2 := 0, 0 //cycles of lenght > 2, cycles of length == 2
	for i := 0; i < n; i++ {
		if !visited[i] {
			l := 0
			for j := i; visited[j] == false; j = favorite[j] {
				visited[j] = true
				l++
			}
			if l == 2 {
				res_2 += 2 + path[i] + path[favorite[i]]
			} else {
				res = max(res, l)
			}
		}
	}
	return max(res, res_2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func make_graph(edges []int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i]] = append(g[edges[i]], i)
	}
	return g
}

func main() {
	fmt.Println(maximum_employees_to_invited_to_a_meeting([]int{2, 2, 1, 2}))
}
