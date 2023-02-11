package main

import "fmt"

/*
	Union find : Here I am using
	path compression as this is the
	most efficient way of doing union
	find.
	Path compression can be done in
	two ways, one using by size of
	the tree and other by using height.
	The time complexity of each operation
	is O(lgn)
*/

//this function returns the parent of each node
func unionFind(edges [][]int, n int) []int {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	for i := 0; i < len(edges); i++ {
		if find(edges[i][0], &parent) != find(edges[i][1], &parent) {
			union(edges[i][0], edges[i][1], &parent, &size)
		}
	}
	return parent
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
	edges := [][]int{
		{0, 1}, {1, 2}, {3, 4},
	}
	fmt.Println(unionFind(edges, 5))
}
