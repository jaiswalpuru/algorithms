package main

import (
	"container/heap"
	"fmt"
)

/*
	Kruskal algorithm :
	find minimum spanning forest, if graph is connected
	it returns the minimum spanning tree.
	O(ElogV)
*/

type Edge struct{ u, v, wt int }

type I []Edge

func (m I) Len() int              { return len(m) }
func (m I) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m I) Less(i, j int) bool    { return m[i].wt < m[j].wt }
func (m *I) Push(val interface{}) { *m = append(*m, val.(Edge)) }
func (m *I) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func kruskal(edges []Edge, n int) int {
	//push into priority queue
	mh := &I{}
	for i := 0; i < len(edges); i++ {
		heap.Push(mh, edges[i])
	}
	//initialize parent and size
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	ed := []Edge{}
	min_cost := 0
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Edge)
		u, v, wt := curr.u, curr.v, curr.wt
		if find(u, &parent) == find(v, &parent) {
			continue
		}
		union(u, v, &parent, &size)
		min_cost += wt
		ed = append(ed, Edge{u, v, wt})
		//stop if all the nodes are connected
		if size[parent[0]] == n {
			break
		}
	}
	if size[parent[0]] != n {
		fmt.Println("MST does not exist")
		return -1
	}
	fmt.Println(ed)
	return min_cost
}

func find(x int, parent *[]int) int {
	if x != (*parent)[x] {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] >= (*size)[y] {
		(*size)[x] += (*size)[y]
		(*parent)[y] = x
	} else {
		(*size)[y] += (*size)[x]
		(*parent)[x] = y
	}
}

func main() {
	g1 := []Edge{
		{0, 1, 5}, {1, 2, 4}, {2, 9, 2},
		{0, 4, 1}, {0, 3, 4}, {1, 3, 2},
		{2, 7, 4}, {2, 8, 1}, {9, 8, 0},
		{4, 5, 1}, {5, 6, 7}, {6, 8, 4},
		{4, 3, 2}, {5, 3, 5}, {3, 6, 11},
		{6, 7, 1}, {3, 7, 2}, {7, 8, 6},
	}
	fmt.Println(kruskal(g1, 10))
}
