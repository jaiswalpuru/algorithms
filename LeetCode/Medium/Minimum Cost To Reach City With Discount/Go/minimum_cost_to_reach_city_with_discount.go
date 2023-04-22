package main

import (
	"container/heap"
	"fmt"
	"math"
)

type P struct {
	v,
	w int
}

type I struct {
	vertex int
	dis    int
	cost   int
}

type Q []I

func (mh Q) Len() int              { return len(mh) }
func (mh Q) Less(i, j int) bool    { return mh[i].cost < mh[j].cost }
func (mh Q) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *Q) Push(val interface{}) { *mh = append(*mh, val.(I)) }
func (mh *Q) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func make_graph(edges [][]int) map[int][]P {
	n := len(edges)
	graph := make(map[int][]P)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], P{edges[i][1], edges[i][2]})
		graph[edges[i][1]] = append(graph[edges[i][1]], P{edges[i][0], edges[i][2]})
	}
	return graph
}

func minimum_cost_to_reach_city_without_discount(n int, edges [][]int, discount int) int {
	graph := make_graph(edges)
	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, discount+1)
		for j := 0; j < discount+1; j++ {
			distance[i][j] = math.MaxInt64
		}
	}

	for i := 0; i <= discount; i++ {
		distance[0][i] = 0
	}

	mh := &Q{}
	heap.Push(mh, I{vertex: 0, dis: discount, cost: 0})

	for mh.Len() > 0 {
		curr := heap.Pop(mh).(I)
		u, d, cost := curr.vertex, curr.dis, curr.cost

		if u == n-1 {
			return cost
		}

		for i := 0; i < len(graph[u]); i++ {
			//moving forward without discount
			if distance[graph[u][i].v][d] > graph[u][i].w+cost {
				distance[graph[u][i].v][d] = graph[u][i].w + cost
				heap.Push(mh, I{vertex: graph[u][i].v, dis: d, cost: distance[graph[u][i].v][d]})
			}

			//taking discount into consideration
			if d > 0 && distance[graph[u][i].v][d-1] > graph[u][i].w/2+cost {
				distance[graph[u][i].v][d-1] = graph[u][i].w/2 + cost
				heap.Push(mh, I{vertex: graph[u][i].v, dis: d - 1, cost: distance[graph[u][i].v][d-1]})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimum_cost_to_reach_city_without_discount(5, [][]int{
		{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 3}, {3, 4, 2},
	}, 1))
}
