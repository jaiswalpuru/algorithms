package main

import (
	"fmt"
	"sort"
)

func find_all_people_with_secret(n int, meetings [][]int, first_person int) []int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	g := make_graph(meetings)
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	union(0, first_person, &parent, &size)
	visited := make(map[int]bool)
	set := []int{}
	for i := 0; i < len(meetings); i++ {
		set = []int{}
		t := meetings[i][2]
		if visited[t] {
			continue
		}
		visited[t] = true
		meet := g[t]
		for j := 0; j < len(meet); j++ {
			union(meet[j].x, meet[j].y, &parent, &size)
			set = append(set, meet[j].x)
			set = append(set, meet[j].y)
		}
		for j := 0; j < len(set); j++ {
			if !(find(0, &parent) == find(set[j], &parent)) {
				parent[set[j]] = set[j]
			}
		}
	}
	secrets := []int{}
	for i := 0; i < n; i++ {
		if find(0, &parent) == find(i, &parent) {
			secrets = append(secrets, i)
		}
	}
	return secrets
}

func union(u, v int, parent *[]int, size *[]int) {
	u = find(u, parent)
	v = find(v, parent)
	if u == v {
		return
	}
	if (*size)[u] > (*size)[v] {
		(*parent)[v] = u
		(*size)[u] += (*size)[v]
	} else {
		(*parent)[u] = v
		(*size)[v] += (*size)[u]
	}
}

func find(u int, parent *[]int) int {
	if u == (*parent)[u] {
		return u
	}
	(*parent)[u] = find((*parent)[u], parent)
	return (*parent)[u]
}

type Pair struct{ x, y int }

func make_graph(edges [][]int) map[int][]Pair {
	g := make(map[int][]Pair)
	n := len(edges)
	for i := 0; i < n; i++ {
		g[edges[i][2]] = append(g[edges[i][2]], Pair{edges[i][0], edges[i][1]})
	}
	return g
}

func main() {
	fmt.Println(find_all_people_with_secret(6, [][]int{
		{1, 2, 5}, {2, 3, 8}, {1, 5, 10},
	}, 1))
}
