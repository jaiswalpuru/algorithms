package main

import "fmt"

func process_requested_friend_requests(n int, restrictions [][]int, requests [][]int) []bool {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	res := make([]bool, len(requests))
	for i := 0; i < len(requests); i++ {
		flag := true
		x, y := find(requests[i][0], &parent), find(requests[i][1], &parent)
		if x != y {
			for i := 0; i < len(restrictions); i++ {
				r1, r2 := find(restrictions[i][0], &parent), find(restrictions[i][1], &parent)
				if (r1 == x && r2 == y) || (r1 == y && r2 == x) {
					flag = false
					break
				}
			}
		}
		res[i] = flag
		if flag {
			union(requests[i][0], requests[i][1], &parent, &size)
		}
	}
	return res
}

func find(x int, parent *[]int) int {
	if x == (*parent)[x] {
		return x
	}
	(*parent)[x] = find((*parent)[x], parent)
	return (*parent)[x]
}

func union(x, y int, parent, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] > (*size)[y] {
		(*parent)[y] = x
		(*size)[x] += (*size)[y]
	} else {
		(*parent)[x] = y
		(*size)[y] += (*size)[x]
	}
}

func main() {
	fmt.Println(process_requested_friend_requests(3, [][]int{{0, 1}}, [][]int{{0, 2}, {2, 1}}))
}
