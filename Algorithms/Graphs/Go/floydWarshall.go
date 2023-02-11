package main

import (
	"fmt"
	"math"
)

//source watch william fiset and read CLRS

//floyd warshall algorithm - for all pair shortest path.
//this means it can find the shortest path from all the pairs.
//O(V^3) which is ideal for graphs with no largest than a
// couple of 100 nodes.

func floydWarshall(graph [][]float64, n int) [][]float64 {
	//the next matrix will be used for path reconstruction
	// dis matrix is distance matrix which is deep copy of graph
	dis := make([][]float64, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]float64, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = graph[i][j]
			if graph[i][j] != math.Inf(1) {
				next[i][j] = j //if the path from i to j is not +inf then
				//the next node which we can go from i is j
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dis[i][j] > dis[i][k]+dis[k][j] {
					dis[i][j] = dis[i][k] + dis[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}
	//detect -ve cycle
	// to check for negative re run the FW algo second time
	// if this time the distance can be improved then we can say
	// that it has negative cycle.
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dis[i][j] > dis[i][k]+dis[k][j] {
					dis[i][j] = math.Inf(-1)
					next[i][j] = -1
				}
			}
		}
	}
	fmt.Println("Next array : ", next)
	return dis
}

func makeGraph(n int) [][]float64 {
	var mat [][]float64
	mat = make([][]float64, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			mat[i][j] = math.Inf(1)
			if i == j {
				mat[i][j] = 0
			}
		}
	}
	return mat
}

func main() {
	var n int = 7
	g := makeGraph(n)
	// Add some edge values.
	g[0][1] = 2
	g[0][2] = 5
	g[0][6] = 10
	g[1][2] = 2
	g[1][4] = 11
	g[2][6] = 2
	g[6][5] = 11
	g[4][5] = 1
	g[5][4] = -2
	dis := floydWarshall(g, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("Shortest path from node %d to node %d is %f\n", i, j, dis[i][j])
		}
	}
}
