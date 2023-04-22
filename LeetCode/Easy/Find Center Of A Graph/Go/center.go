package main

import "fmt"

func center(edges [][]int) int {
	if edges[0][0] == edges[1][0] {
		return edges[0][0]
	} else if edges[0][1] == edges[1][0] {
		return edges[0][1]
	} else if edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}

func main() {
	fmt.Println(center([][]int{
		{1, 2}, {2, 3}, {4, 2},
	}))
}
