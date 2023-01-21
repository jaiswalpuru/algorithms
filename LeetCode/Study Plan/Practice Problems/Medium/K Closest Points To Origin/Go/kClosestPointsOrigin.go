package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Pair struct {
	points []int
	dis    float64
}

type M []Pair

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i].dis < m[j].dis }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func euclideanDistance(x, y int) float64 {
	return math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
}

func kClosest(points [][]int, k int) [][]int {
	m := &M{}
	for i := 0; i < len(points); i++ {
		heap.Push(m, Pair{points[i], euclideanDistance(points[i][0], points[i][1])})
	}
	res := make([][]int, k)
	for k > 0 {
		top := heap.Pop(m).(Pair)
		res[k-1] = top.points
		k--
	}
	return res
}

func main() {
	fmt.Println(kClosest([][]int{
		{3, 3}, {5, -1}, {-2, 4},
	}, 2))
}
