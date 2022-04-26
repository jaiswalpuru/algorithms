package main

import "fmt"
import "container/heap"

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func mahattan_distance(a, b []int) int { return abs(a[0]-b[0]) + abs(a[1]-b[1]) }

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

func min_cost_to_connect_all_points(arr [][]int) int {
	n := len(arr)
	visited := make([]bool, n)
	connected, min_cost := 0, 0
	mh := &MinHeap{Point{0, 0, 0}}
	for mh.Len() > 0 && connected < n {
		curr := heap.Pop(mh).(Point)
		src, dst, dis := curr.src, curr.dst, curr.dis
		if visited[src] && visited[dst] {
			continue
		}

		min_cost += dis
		i := dst
		visited[i] = true
		connected++
		for j := 0; j < n; j++ {
			if visited[j] {
				continue
			}
			cost := mahattan_distance(arr[i], arr[j])
			heap.Push(mh, Point{i, j, cost})
		}
	}
	return min_cost
}

func main() {
	fmt.Println(min_cost_to_connect_all_points([][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}))
}
