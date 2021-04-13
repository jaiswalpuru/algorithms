package main

import (
	"container/heap"
	"fmt"
)

type Vertex int

type PriorityQueue struct {
	items []Vertex
	m     map[Vertex]int // value to index
	pr    map[Vertex]int //value to priority
}

func (pq *PriorityQueue) Len() int           { return len(pq.items) }
func (pq *PriorityQueue) Less(i, j int) bool { return pq.pr[pq.items[i]] < pq.pr[pq.items[j]] }
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i]] = i
	pq.m[pq.items[j]] = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(Vertex)
	pq.m[item] = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(pq.items)
	item := pq.items[n-1]
	pq.m[item] = -1
	pq.items = old[:n-1]
	return item
}

func NewGraph(ids map[string]Vertex) AdjacencyMap {
	G := AdjacencyMap{ids: ids}
	G.names = make(map[Vertex]string)
	for k, v := range ids {
		G.names[v] = k
	}
	G.edges = make(map[Vertex]map[Vertex]int)
	return G
}

//update modifies the priority of an item in the queue
func (pq *PriorityQueue) update(item Vertex, priority int) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}

func (pq *PriorityQueue) addWithPriority(item Vertex, priority int) {
	heap.Push(pq, item)
	pq.update(item, priority)
}

const (
	Infinity      = int(^uint(0) >> 1) //maximum integer
	Uninitialized = -1                 //indicates not yet processed vertex
)

type Graph interface {
	Vertices() []Vertex
	Neighbors(v Vertex) []Vertex
	Weight(u, v Vertex) int
}

//AdjacencyMap is a graph of strings that satisfies the Graph interface{}
type AdjacencyMap struct {
	ids   map[string]Vertex
	names map[Vertex]string
	edges map[Vertex]map[Vertex]int
}

func (G AdjacencyMap) AddEdge(u, v string, w int) {
	if _, ok := G.edges[G.ids[u]]; !ok {
		G.edges[G.ids[u]] = make(map[Vertex]int)
	}
	G.edges[G.ids[u]][G.ids[v]] = w
}

func (G AdjacencyMap) GetPath(v Vertex, previous map[Vertex]Vertex) (path string) {
	path = G.names[v]
	for previous[v] >= 0 {
		v = previous[v]
		path = G.names[v] + path
	}
	return path
}

func (G AdjacencyMap) Vertices() (vertices []Vertex) {
	for _, v := range G.ids {
		vertices = append(vertices, v)
	}
	return vertices
}

func (G AdjacencyMap) Neighbors(u Vertex) (vertices []Vertex) {
	for v := range G.edges[u] {
		vertices = append(vertices, v)
	}
	return vertices
}

func (G AdjacencyMap) Weight(u, v Vertex) int { return G.edges[u][v] }

func Dijkstra(G Graph, source Vertex) (distance map[Vertex]int, previous map[Vertex]Vertex) {
	distance = make(map[Vertex]int)
	previous = make(map[Vertex]Vertex)
	distance[source] = 0
	q := &PriorityQueue{[]Vertex{}, make(map[Vertex]int), make(map[Vertex]int)}
	for _, v := range G.Vertices() {
		if v != source {
			distance[v] = Infinity
		}
		previous[v] = Uninitialized
		q.addWithPriority(v, distance[v])
	}
	for len(q.items) != 0 {
		u := heap.Pop(q).(Vertex)
		for _, v := range G.Neighbors(u) {
			alt := distance[u] + G.Weight(u, v)
			if alt < distance[v] {
				distance[v] = alt
				previous[v] = u
				q.update(v, alt)
			}
		}
	}
	return distance, previous
}

func main() {
	G := NewGraph(map[string]Vertex{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	})
	G.AddEdge("a", "b", 7)
	G.AddEdge("a", "c", 9)
	G.AddEdge("a", "f", 10)
	G.AddEdge("b", "c", 10)
	G.AddEdge("c", "d", 11)
	G.AddEdge("c", "f", 2)
	G.AddEdge("d", "e", 6)
	G.AddEdge("e", "f", 9)
	distance, previous := Dijkstra(G, G.ids["a"])
	fmt.Printf("Distance to %s is %d, Path : %s\n", "e", distance[G.ids["e"]], G.GetPath(G.ids["e"], previous))

}
