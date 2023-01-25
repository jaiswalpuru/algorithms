package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Pair struct{ v, wt int }
type M []Pair

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i].wt < m[j].wt }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

// Johnson algorithm for all pair shortest path
// O(V^2logV + VE). It is better than floyd warshall
// which takes O(V^3)
// this algo uses both dijkstra and bellman as subroutines.

//In johnson first we add a pseudo vertex and calculate the
// distance matrix from that point, that is reweight the vertex
// and after that remove the pseudo vertex and apply dijkstra.
func johnsons(edges [][]int, n int) {
	//add pseudo vertex to all the other vertices
	pseudoVertex := n
	for i := 0; i < n; i++ {
		edges = append(edges, []int{pseudoVertex, i, 0})
	}
	fmt.Println("After adding the pseudo vertex : ", edges)
	totalNode := n + 1
	dis := make([]float64, totalNode)
	for i := 0; i < totalNode; i++ {
		dis[i] = math.Inf(1)
	}
	dis[pseudoVertex] = 0
	//run bellman ford to get the h value (dis from source)
	//basically reweighting
	for i := 0; i < totalNode-1; i++ {
		for j := 0; j < len(edges); j++ {
			u, v, wt := edges[j][0], edges[j][1], float64(edges[j][2])
			if dis[v] > dis[u]+wt {
				dis[v] = dis[u] + wt
			}
		}
	}

	//running the bellman ford alogrithm again to check for negative wt cycles.
	for i := 0; i < totalNode-1; i++ {
		for j := 0; j < len(edges); j++ {
			u, v, wt := edges[j][0], edges[j][1], float64(edges[j][2])
			if dis[v] > dis[u]+wt {
				fmt.Println("the graph has negative wt cycle")
				return
			}
		}
	}

	fmt.Println("After running the bellman the distance vector : ", dis)

	//we need to remove the edges which was introduced
	newEdges := [][]int{}
	for i := 0; i < totalNode; i++ {
		if edges[i][0] != pseudoVertex {
			newEdges = append(newEdges, edges[i])
		}
	}
	fmt.Println("after removing the vertex which was introduced ", newEdges)

	//reweighting the edges
	for i := 0; i < len(newEdges); i++ {
		u, v := newEdges[i][0], newEdges[i][1]
		newEdges[i][2] = int(float64(newEdges[i][2]) + dis[u] - dis[v])
	}
	fmt.Println("after reweighting the edges : ", newEdges)

	//converting the graph into adjacency list
	g := makeGraph(edges)

	//run dijkstra for every vertex as the source
	for u := 0; u < n; u++ {
		newDis, _ := dijkstra(u, g, n)
		for v := 0; v < n; v++ {
			newDis[v] = newDis[v] + dis[v] - dis[u]
		}
		fmt.Println("for vertex ", u, " distance to other node : ", newDis)
		fmt.Println()
	}
}

func dijkstra(src int, g map[int][]Pair, n int) ([]float64, []int) {
	dis := make([]float64, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.Inf(1)
	}
	dis[src] = 0
	parent[src] = -1
	visited := make(map[int]bool)
	mh := &M{}
	heap.Push(mh, Pair{src, 0})
	for mh.Len() > 0 {
		currElement := heap.Pop(mh).(Pair)
		u, wt := currElement.v, currElement.wt
		if visited[u] {
			continue
		}
		visited[u] = true
		for i := 0; i < len(g[u]); i++ {
			v, vwt := g[u][i].v, g[u][i].wt
			fmt.Println(wt, vwt, float64(wt+vwt), u, v)
			if !visited[v] {
				if dis[v] > float64(wt+vwt) {
					parent[v] = u
					dis[v] = float64(wt + vwt)
					heap.Push(mh, Pair{v, int(dis[v])})
				}
			}
		}
	}
	return dis, parent
}

func makeGraph(edges [][]int) map[int][]Pair {
	g := make(map[int][]Pair)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], Pair{edges[i][1], edges[i][2]})
	}
	return g
}

func main() {
	edges := [][]int{
		{0, 1, -5}, {0, 3, 3}, {0, 2, 2},
		{1, 2, 4}, {2, 3, 1},
	}
	johnsons(edges, 4)
}
