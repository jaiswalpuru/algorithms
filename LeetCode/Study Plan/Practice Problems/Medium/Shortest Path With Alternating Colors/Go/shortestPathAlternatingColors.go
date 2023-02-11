package main

import (
	"fmt"
)

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	res[0] = 0
	q := []Pair{{0, -1}}
	dis := 0
	graph := makeGraph(redEdges, blueEdges)
	visited := make(map[[2]int]bool)

	for len(q) > 0 {
		qSize := len(q)
		for i := 0; i < qSize; i++ {
			curr := q[0]
			q = q[1:]
			if visited[[2]int{curr.val, curr.color}] {
				continue
			}
			if res[curr.val] != -1 {
				res[curr.val] = min(dis, res[curr.val])
			} else {
				res[curr.val] = dis
			}

			visited[[2]int{curr.val, curr.color}] = true
			for j := 0; j < len(graph[curr.val]); j++ {
				v := graph[curr.val][j]
				if !visited[[2]int{v.val, v.color}] {
					if curr.color == -1 {
						q = append(q, Pair{v.val, v.color})
					} else {
						if curr.color != v.color {
							q = append(q, Pair{v.val, v.color})
						}
					}
				}
			}
		}
		dis++
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type Pair struct {
	val   int
	color int
}

func makeGraph(redEdges [][]int, blueEdges [][]int) map[int][]Pair {
	g := make(map[int][]Pair)
	for i := 0; i < len(redEdges); i++ {
		g[redEdges[i][0]] = append(g[redEdges[i][0]], Pair{redEdges[i][1], 0})
	}
	for i := 0; i < len(blueEdges); i++ {
		g[blueEdges[i][0]] = append(g[blueEdges[i][0]], Pair{blueEdges[i][1], 1})
	}
	return g
}

func main() {
	fmt.Println(shortestAlternatingPaths(3,
		[][]int{
			{0, 1}, {0, 2},
		},
		[][]int{
			{1, 0},
		}))
}
