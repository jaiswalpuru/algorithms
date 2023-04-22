package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

type Pair struct {
	val, d int
}

//--------------------brute force-------------------------
func sum_of_distances_in_tree(n int, edges [][]int) []int {
	sum := make([]int, n)
	g := make_graph(edges)
	for i := 0; i < n; i++ {
		s := 0
		visited := make(map[int]bool)
		q := []Pair{{val: i, d: 0}}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			visited[curr.val] = true
			s += curr.d
			for j := 0; j < len(g[curr.val]); j++ {
				if !visited[g[curr.val][j]] {
					q = append(q, Pair{val: g[curr.val][j], d: curr.d + 1})
				}
			}
		}
		sum[i] = s
	}
	return sum
}

//--------------------brute force-------------------------

//--------------------efficient-------------------------
func sum_of_distances_in_tree_eff(n int, edges [][]int) []int {
	g := make_graph(edges)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = 1
	}
	sum := make([]int, n)
	dfs_one(g, 0, -1, &count, &sum)
	dfs_two(g, 0, -1, &count, &sum, n)
	return sum
}

func dfs_one(g map[int][]int, node, parent int, cnt *[]int, sum *[]int) {
	for i := 0; i < len(g[node]); i++ {
		if g[node][i] != parent {
			dfs_one(g, g[node][i], node, cnt, sum)
			(*cnt)[node] += (*cnt)[g[node][i]]
			(*sum)[node] += (*sum)[g[node][i]] + (*cnt)[g[node][i]]
		}
	}
}

func dfs_two(g map[int][]int, node, parent int, cnt *[]int, sum *[]int, n int) {
	for i := 0; i < len(g[node]); i++ {
		if g[node][i] != parent {
			(*sum)[g[node][i]] = (*sum)[node] - (*cnt)[g[node][i]] + n - (*cnt)[g[node][i]]
			dfs_two(g, g[node][i], node, cnt, sum, n)
		}
	}
}

//--------------------efficient-------------------------

func main() {
	fmt.Println(sum_of_distances_in_tree_eff(6, [][]int{
		{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5},
	}))
}
