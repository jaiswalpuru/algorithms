package main

import (
	"fmt"
	"sort"
)

func account_merge(acc [][]string) [][]string {
	graph := make(map[string][]string)
	visited := make(map[string]bool)

	n := len(acc)

	for i := 0; i < n; i++ {
		first_email := acc[i][1]

		m := len(acc[i])
		for j := 2; j < m; j++ {
			email := acc[i][j]
			graph[email] = append(graph[email], first_email)
			graph[first_email] = append(graph[first_email], email)
		}
	}

	res := [][]string{}
	for i := 0; i < n; i++ {
		name := acc[i][0]
		first_email := acc[i][1]
		if !visited[first_email] {
			temp := []string{name}
			dfs(graph, first_email, &visited, &temp)
			sort.Strings(temp[1:])
			res = append(res, temp)
		}
	}
	return res
}

func dfs(graph map[string][]string, start_email string, visited *map[string]bool, temp *[]string) {
	(*visited)[start_email] = true
	*temp = append(*temp, start_email)
	if _, ok := graph[start_email]; !ok {
		return
	}
	for i := 0; i < len(graph[start_email]); i++ {
		if !(*visited)[graph[start_email][i]] {
			dfs(graph, graph[start_email][i], visited, temp)
		}
	}
}

func main() {
	fmt.Println(account_merge([][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}))
}
