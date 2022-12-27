package main

import (
	"fmt"
	"sort"
)

func find(x int, parent *[]int) int {
	if x == (*parent)[x] {
		return x
	}
	(*parent)[x] = find((*parent)[x], parent)
	return (*parent)[x]
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] >= (*size)[y] {
		(*parent)[y] = x
		(*size)[x] += (*size)[y]
	} else {
		(*parent)[x] = y
		(*size)[y] += (*size)[x]
	}
}

func earliest_moment_when_everyone_becomes_friends(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	latest_time := logs[0][0]
	for i := 0; i < len(logs); i++ {
		time := logs[i][0]
		x, y := logs[i][1], logs[i][2]
		if find(x, &parent) == find(y, &parent) {
			continue
		} else {
			latest_time = time
			union(x, y, &parent, &size)
		}
	}
	for i := 0; i < n; i++ {
		if find(0, &parent) != find(i, &parent) {
			return -1
		}
	}
	return latest_time
}

func main() {
	fmt.Println(earliest_moment_when_everyone_becomes_friends([][]int{
		{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {20190224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5},
	}, 6))
}
