package main

import (
	"bufio"
	"fmt"
	"os"
)

type GraphType string

const mod = 1000000007 //10^9 modulo
const DIRECTED = "DIRECTED"
const UNDIRECTED = "UNDIRECTED"

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

type trie struct {
	letter   rune
	children []*trie
	isLeaf   bool
	metaData map[string]string
}

func newTrie() *trie {
	nT := &trie{}
	nT.children = []*trie{}
	nT.letter = '#' //Initital value say works as a root node
	nT.metaData = make(map[string]string)
	return nT
}

type Pair struct {
	a, b int
}

//sort pair for struct based on the first value
type Pairs []Pair

func (a Pairs) Len() int           { return len(a) }
func (a Pairs) Less(i, j int) bool { return a[i].a < a[j].a }
func (a Pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func gcd(a, b int64) int64 {
	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func reverseRune(s *[]rune) {
	for i, j := 0, len(*s)-1; i < len(*s)/2; {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i++
		j--
	}
}

func pushBack(val int, arr *[]int) {
	*arr = append(*arr, val)
}

func popBack(arr *[]int) int {
	val := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return val
}

func addEdge(u, v int, graph *[][]int, gType GraphType) {
	if gType == DIRECTED {
		(*graph)[u] = append((*graph)[u], v)
	} else {
		(*graph)[u] = append((*graph)[u], v)
		(*graph)[v] = append((*graph)[v], u)
	}
}

func addEdgeWeigts(u int, v Pair, graph *[][]Pair, gType GraphType) {
	if gType == DIRECTED {
		(*graph)[u] = append((*graph)[u], v)
	} else {
		(*graph)[u] = append((*graph)[u], v)
		(*graph)[v.a] = append((*graph)[v.a], Pair{a: u, b: v.b})
	}
}

//shortest path in unweighted graph from the source vertex
//s - > source node, nodes -> total number of nodes, distace 1-D array with -1 initally assigned to it
func unweightedShortestPath(graph [][]int, s int, distance []int) {
	queue := []int{s}
	for i := 0; i < len(graph); i++ {
		distance[i] = -1
	}
	//We can also keep an array for storing the path of the adjacent nodes
	distance[s] = 0 //distance of source to source is 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for i := 0; i < len(graph[v]); i++ {
			if distance[graph[v][i]] == -1 {
				distance[graph[v][i]] += distance[v] + 1
				queue = append(queue, graph[v][i])
			}
		}
	}
}

//argument is [i][j]int in which i represents the current node and arr[i] will contain a array of values through which it is connected
func dfsUtil(curr int, graph [][]int, visited *[]bool, ordering *[]int) {
	if !(*visited)[curr] {
		(*visited)[curr] = true
		*ordering = append(*ordering, curr)
		for j := 0; j < len(graph[curr]); j++ {
			if !(*visited)[graph[curr][j]] {
				dfsUtil(graph[curr][j], graph, visited, ordering)
			}
		}
	}
}

//argument is [i][j]int in which i represents the current node and arr[i] will contain a array of values through which it is connected
func bfsUtil(val int, graph [][]int, visited *[]bool, ordering *[]int) {
	queue := []int{}
	queue = append(queue, val)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if (*visited)[curr] {
			continue
		}
		(*ordering) = append((*ordering), curr)
		(*visited)[curr] = true
		for j := 0; j < len(graph[curr]); j++ {
			if (*visited)[graph[curr][j]] {
				continue
			}
			queue = append(queue, graph[curr][j])
		}
	}
}

/*
	Start your code from here, above are helper functions
*/

//define constants here
const ()

//define global variables here
var ()

func main() {
	defer writer.Flush()
}