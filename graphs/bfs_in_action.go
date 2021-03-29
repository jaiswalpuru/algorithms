package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"os"
)

type GraphType string

const (
	DIRECTED   = "DIRECTED"
	UNDIRECTED = "UNDIRECTED"
)

type Pair struct {
	a, b int
}

type PairWithDistance struct {
	a, b, Distance int
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var scanner *bufio.Scanner = bufio.NewScanner(reader)

func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }
func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = (n * 10) + int(v-'0')
	}
	return
}

type pairs []Pair

func (a pairs) Len() int           { return len(a) }
func (a pairs) Less(i, j int) bool { return a[i].a < a[j].a }
func (a pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Node struct {
	Next   *Node
	Weight int
	Key    int
}

type AdjacencyList struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjList   []*Node
}

func (G *AdjacencyList) Init() {
	G.AdjList = make([]*Node, G.Vertices)
	G.Edges = 0
	for i := 0; i < G.Vertices; i++ {
		G.AdjList[i] = &Node{}
	}
}

//Recursive method to add node to the last available slot
func (node Node) AddNode(value int) *Node {
	n := node.Next
	if n == nil {
		newNode := &Node{Next: &Node{}, Key: value}
		return newNode
	}
	nd := n.AddNode(value)
	node.Next = nd
	return &node
}

func (node Node) FindNextNode(key int) (*Node, error) {
	n := node
	if n == (Node{}) {
		return &Node{}, errors.New("Node not found")
	}
	if n.Key == key {
		return &n, nil
	}
	nd := n.Next
	return nd.FindNextNode(key)
}

func (G *AdjacencyList) AddEdge(vertexOne, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bound")
	}
	node := G.AdjList[vertexOne].AddNode(vertexTwo)
	G.AdjList[vertexOne] = node
	G.Edges++
	if G.GraphType == UNDIRECTED {
		node := G.AdjList[vertexTwo].AddNode(vertexOne)
		G.AdjList[vertexTwo] = node
		G.Edges++
	}
	return nil
}

func (G *AdjacencyList) HasEdge(vertexOne, vertexTwo int) bool {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}
	nodeAdj := G.AdjList[vertexOne]
	if nodeAdj == (&Node{}) {
		return false
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return true
	}
	return false
}

func (G *AdjacencyList) GetAdjacencyNodesForVertex(vertex int) map[int]bool {
	if vertex >= G.Vertices || vertex < 0 {
		return map[int]bool{}
	}
	nodeAdj := G.AdjList[vertex]
	nextNode := nodeAdj
	nodes := map[int]bool{}
	for nextNode != (&Node{}) && nextNode != nil {
		nodes[nextNode.Key] = true
		nextNode = nextNode.Next
	}
	return nodes
}

func BFSWrapper(G *AdjacencyList) {
	visited := map[int]bool{}
	distance := make(map[int]int)
	bfs(G, visited, 0, &distance)
	fmt.Println(distance)
}

func bfs(G *AdjacencyList, visited map[int]bool, node int, mp *map[int]int) {
	queue := make([]int, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current] {
			continue
		}
		visited[current] = true
		adjNode := G.GetAdjacencyNodesForVertex(current)
		for key := range adjNode {
			if key != 0 {
				queue = append(queue, key)
				if (*mp)[key] == 0 {
					(*mp)[key] = (*mp)[current] + 1
				}
			}
		}
	}
}

// else {
// 	if current != 0 {
// 		fmt.Println(current)
// 		(*mp)[current] = (*mp)[key] + 1
// 	}
// }

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for t > 0 {
		var n, m, s, k, u, v, x int
		scanf("%d %d %d %d\n", &n, &m, &s, &k)
		l := make([]list.List, n+1)

		for i := 0; i < m; i++ {
			scanf("%d %d\n", &u, &v)
			l[u].PushBack(v)
			l[v].PushBack(u)
		}

		building := make([]list.List, n+1)
		for i := 0; i < s; i++ {
			if i == n-1 {
				scanf("%d\n", &x)
			} else {
				scanf("%d ", &x)
			}
			building[x].PushBack(i)
		}

		for i := 0; i < len(building); i++ {
			fmt.Print(i, "->")
			for e := building[i].Front(); e != nil; e = e.Next() {
				fmt.Print(e.Value.(int))
			}
			fmt.Println()
		}

		fmt.Println(building[1].Len())

		vis := make([]int, n+1)
		for i := 0; i <= n; i++ {
			vis[i] = -1
		}
		vis[0] = 0

		queue := make([]int, 0)
		queue = append(queue, 0)
		ans := 0
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			fmt.Println("Len", building[current].Len(), current)
			for (building[current].Len() != 0) && (k != 0) {
				k--
				ans += 2 * vis[current]
				building[current].Remove(building[current].Back())
			}

			for val := l[current].Front(); val != nil; val = val.Next() {
				if vis[val.Value.(int)] > 0 {
					continue
				}
				vis[val.Value.(int)] = vis[current] + 1
				queue = append(queue, val.Value.(int))
			}
		}

		fmt.Println(ans)

		t--
	}
}
