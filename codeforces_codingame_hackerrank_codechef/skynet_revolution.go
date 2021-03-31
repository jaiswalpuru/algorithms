package main

import (
	"fmt"
	"os"
)

func addEdge(u, v int, graph *[][]int) {
	(*graph)[u] = append((*graph)[u], v)
	(*graph)[v] = append((*graph)[v], u)
}

func bfsUtil(val int, g [][]int, v *[]bool, o *[]int) {
	curr := val
	(*v)[curr] = true
	for i := 0; i < len(g[curr]); i++ {
		if (*v)[g[curr][i]] {
			continue
		}
		(*o) = append((*o), g[curr][i])
	}
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)
	fmt.Fprintln(os.Stderr, N, L, E)
	graph := make([][]int, N)

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)
		fmt.Fprintln(os.Stderr, N1, N2)
		addEdge(N1, N2, &graph)
	}
	indexExit := make(map[int]bool)
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		indexExit[EI] = true
	}
	fmt.Fprintln(os.Stderr, indexExit)
	for {
		ordering := []int{}
		visited := make([]bool, N)
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)
		bfsUtil(SI, graph, &visited, &ordering)
		fmt.Fprintln(os.Stderr, SI, ordering)
		f := false
		for i := 0; i < len(ordering); i++ {
			if indexExit[ordering[i]] == true {
				f = true
				fmt.Println(SI, ordering[i])
				break
			}
		}

		if f == false {
			fmt.Println(SI, ordering[0])
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// Example: 0 1 are the indices of the nodes you wish to sever the link between
	}
}
