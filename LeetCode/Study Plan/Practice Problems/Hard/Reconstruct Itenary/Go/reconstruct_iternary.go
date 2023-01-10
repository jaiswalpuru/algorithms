package main

import (
	"fmt"
	"sort"
)

//-----------------------using the concept of eulerian cycle-------------------------
func reconstruct_iternary_euler_path_concept(tickets [][]string) []string {
	//make graph
	g := make(map[string][]string)
	for i := 0; i < len(tickets); i++ {
		g[tickets[i][0]] = append(g[tickets[i][0]], tickets[i][1])
	}
	for k, _ := range g {
		sort.Strings(g[k])
	}
	res := []string{}
	dfs("JFK", g, &res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func dfs(src string, g map[string][]string, res *[]string) {
	if len(g[src]) != 0 {
		for len(g[src]) > 0 {
			curr := g[src][0]
			g[src] = g[src][1:]
			dfs(curr, g, res)
		}
	}
	*res = append(*res, src)
}

//-----------------------using the concept of eulerian cycle-------------------------

func reconstruct_iternary(tickets [][]string) []string {
	//make graph
	g := make(map[string][]string)
	for i := 0; i < len(tickets); i++ {
		g[tickets[i][0]] = append(g[tickets[i][0]], tickets[i][1])
	}
	visited := make(map[string][]bool)
	for k, v := range g {
		sort.Strings(g[k])
		visited[k] = make([]bool, len(v))
	}
	res := []string{}
	backtrack("JFK", g, &visited, &res, []string{"JFK"}, len(tickets))
	return res
}

func backtrack(src string, g map[string][]string, visited *map[string][]bool, res *[]string, set []string, tickets int) bool {
	if len(set) == tickets+1 {
		*res = append([]string{}, set...)
		return true
	}
	if _, ok := g[src]; !ok {
		return false
	}
	for i := 0; i < len(g[src]); i++ {
		if !(*visited)[src][i] {
			(*visited)[src][i] = true
			set = append(set, g[src][i])
			path := backtrack(g[src][i], g, visited, res, set, tickets)
			set = set[:len(set)-1]
			(*visited)[src][i] = false
			if path {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(reconstruct_iternary_euler_path_concept([][]string{
		{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"},
	}))
}
