package main

import (
	"fmt"
	"math"
	"sort"
)

func find_critical_and_pseudo_critical_edges(n int, edges [][]int) [][]int {
	for i := 0; i < len(edges); i++ {
		edges[i] = append(edges[i], i)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	total_wt := mst_excluding_edge(n, -1, edges)
	critical, pcritical := []int{}, []int{}
	for i := 0; i < len(edges); i++ {
		wt_exc_edge := mst_excluding_edge(n, i, edges)
		if wt_exc_edge > total_wt {
			critical = append(critical, edges[i][3])
		} else {
			wt_inc_edges := mst_including_edge(n, i, edges)
			if wt_inc_edges == total_wt {
				pcritical = append(pcritical, edges[i][3])
			}
		}
	}
	return [][]int{critical, pcritical}
}

func mst_excluding_edge(n, edge_ind int, edges [][]int) int {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	total_wt := 0
	for i := 0; i < len(edges); i++ {
		if i == edge_ind {
			continue
		}
		u, v, wt := edges[i][0], edges[i][1], edges[i][2]
		if union(u, v, &parent, &size) {
			total_wt += wt
		}
	}
	p := find(0, &parent)
	for i := 1; i < n; i++ {
		if find(i, &parent) != p {
			return math.MaxInt64
		}
	}
	return total_wt
}

func mst_including_edge(n, edge_ind int, edges [][]int) int {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	u, v, wt := edges[edge_ind][0], edges[edge_ind][1], edges[edge_ind][2]
	total_wt := wt
	union(u, v, &parent, &size)
	for i := 0; i < len(edges); i++ {
		if edge_ind == i {
			continue
		}
		u, v, wt := edges[i][0], edges[i][1], edges[i][2]
		if union(u, v, &parent, &size) {
			total_wt += wt
		}
	}
	p := find(0, &parent)
	for i := 1; i < n; i++ {
		if find(i, &parent) != p {
			return math.MaxInt64
		}
	}
	return total_wt
}

func find(a int, parent *[]int) int {
	if a != (*parent)[a] {
		(*parent)[a] = find((*parent)[a], parent)
	}
	return (*parent)[a]
}

func union(a, b int, parent *[]int, size *[]int) bool {
	a = find(a, parent)
	b = find(b, parent)
	if a == b {
		return false
	}
	if (*size)[a] >= (*size)[b] {
		(*size)[a] += (*size)[b]
		(*parent)[b] = a
	} else {
		(*size)[b] += (*size)[a]
		(*parent)[a] = b
	}
	return true
}

func main() {
	fmt.Println(find_critical_and_pseudo_critical_edges(5, [][]int{
		{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6},
	}))
}
