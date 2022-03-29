package main

import (
	"container/heap"
	"fmt"
	"math"
)

type C struct{ x, y int }

type Pair struct {
	coord C
	dis   float64
}

func find_euclidean_distance(x, y float64) float64 { return math.Sqrt(x*x + y*y) }

type P []Pair

func (mh P) Len() int              { return len(mh) }
func (mh P) Less(i, j int) bool    { return mh[i].dis < mh[j].dis }
func (mh P) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *P) Push(val interface{}) { *mh = append(*mh, val.(Pair)) }
func (mh *P) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func k_closest_point(points [][]int, k int) [][]int {
	res := [][]int{}

	mh := &P{}
	n := len(points)
	for i := 0; i < n; i++ {
		heap.Push(mh, Pair{coord: C{points[i][0], points[i][1]}, dis: find_euclidean_distance(float64(points[i][0]), float64(points[i][1]))})
	}

	for k > 0 {
		val := heap.Pop(mh).(Pair)
		res = append(res, []int{val.coord.x, val.coord.y})
		k--
	}

	return res
}

func main() {
	fmt.Println(k_closest_point([][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}, 2))
}
