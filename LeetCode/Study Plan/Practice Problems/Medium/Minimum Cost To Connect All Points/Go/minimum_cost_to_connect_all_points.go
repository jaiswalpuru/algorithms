package main

import (
	"container/heap"
	"fmt"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func manhattan_distance(a, b []int) int { return abs(a[0]-b[0]) + abs(a[1]-b[1]) }

type Point struct{ src, dst, dis int }

type MinHeap []Point

func (mh MinHeap) Len() int              { return len(mh) }
func (mh MinHeap) Less(i, j int) bool    { return mh[i].dis < mh[j].dis }
func (mh MinHeap) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(val interface{}) { *mh = append(*mh, val.(Point)) }
func (mh *MinHeap) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

//-----------this is using prims algorithm--------------
func min_cost_to_connect_all_points(points [][]int) int {
	mh := &MinHeap{}
	for i := 1; i < len(points); i++ {
		dis := manhattan_distance(points[0], points[i])
		heap.Push(mh, Point{0, i, dis})
	}
	visited := make([]bool, len(points))
	visited[0] = true
	cost := 0
	conn := 0
	for mh.Len() > 0 && conn < len(points) {
		curr := heap.Pop(mh).(Point)
		if !visited[curr.src] {
			visited[curr.dst] = true
			cost += curr.dis
			for i := 0; i < len(points); i++ {
				if !visited[i] {
					dis := manhattan_distance(points[curr.dst], points[i])
					heap.Push(mh, Point{curr.dst, i, dis})
				}
			}
		}
	}
	return cost
}

//-----------this is using prims algorithm--------------

//-----------kruskals algorithm--------------
func min_cost_to_connect_all_points_kruskals(points [][]int) int {
	mh := &MinHeap{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dis := manhattan_distance(points[i], points[j])
			heap.Push(mh, Point{i, j, dis})
		}
	}
	parent := make([]int, len(points))
	size := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		parent[i] = i
		size[i] = 1
	}
	n := len(points)
	cost := 0
	for mh.Len() > 0 && n > 0 {
		curr := heap.Pop(mh).(Point)
		if find(curr.src, &parent) != find(curr.dst, &parent) {
			union(curr.src, curr.dst, &parent, &size)
			cost += curr.dis
			n--
		}
	}
	return cost
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
	if (*size)[x] > (*size)[y] {
		(*size)[x] += (*size)[y]
		(*parent)[y] = x
	} else {
		(*size)[y] += (*size)[x]
		(*parent)[x] = y
	}
}

//-----------kruskals algorithm--------------

func main() {
	fmt.Println(min_cost_to_connect_all_points_kruskals([][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}))
}
