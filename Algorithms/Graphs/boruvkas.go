package main

import "fmt"

// Boruvkas algorithm : to find the mst
// For weighted, unidrected graph
// O(ElogV)

func boruvkas(edges [][]int, n int) int {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	mst := [][]int{}
	mstCost := 0
	cheapest := make([]int, n)
	for len(mst) < n-1 {
		//fill all the elements as -1
		fill(&cheapest, -1)
		stop := true
		for i := 0; i < len(edges); i++ {
			u, v, wt := edges[i][0], edges[i][1], edges[i][2]
			up, vp := find(u, &parent), find(v, &parent)
			if up == vp {
				continue
			}
			if cheapest[vp] == -1 || wt < edges[cheapest[vp]][2] {
				cheapest[vp] = i
				stop = false
			}
			if cheapest[up] == -1 || wt < edges[cheapest[up]][2] {
				cheapest[up] = i
				stop = false
			}
		}
		if stop {
			break
		}
		for i := 0; i < n; i++ {
			if cheapest[i] == -1 {
				continue
			}
			edge := edges[cheapest[i]]
			if find(edge[0], &parent) == find(edge[1], &parent) {
				continue
			}
			mst = append(mst, edge)
			mstCost += edge[2]
			union(edge[0], edge[1], &parent, &size)
		}
	}
	if len(mst) != n-1 {
		fmt.Println("mst does not exist")
		return -1
	}
	fmt.Println("edges which make mst : ", mst)
	return mstCost
}

func fill(nums *[]int, val int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] = -1
	}
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] > (*size)[y] {
		(*size)[x] += (*size)[y]
		(*parent)[y] = x
	} else {
		(*size)[y] += (*size)[x]
		(*parent)[x] = y
	}
}

func find(x int, parent *[]int) int {
	if x != (*parent)[x] {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
}

func main() {
	edges := [][]int{
		{0, 1, 10}, {0, 2, 6},
		{0, 3, 5}, {1, 3, 15},
		{2, 3, 4},
	}
	fmt.Println("Min cost : ", boruvkas(edges, 4))
}
