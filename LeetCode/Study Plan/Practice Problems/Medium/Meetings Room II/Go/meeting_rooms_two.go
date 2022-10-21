package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type I []int

func (a I) Len() int              { return len(a) }
func (a I) Swap(i, j int)         { a[i], a[j] = a[j], a[i] }
func (a I) Less(i, j int) bool    { return a[i] < a[j] }
func (a *I) Push(val interface{}) { *a = append(*a, val.(int)) }
func (a *I) Pop() interface{} {
	val := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return val
}

func meeting_rooms_two(arr [][]int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	mh := &I{}
	n := len(arr)
	rooms := 1
	heap.Push(mh, arr[0][1])
	for i := 1; i < n; i++ {
		val := heap.Pop(mh).(int)
		if val > arr[i][0] {
			rooms++
			heap.Push(mh, val)
		}
		heap.Push(mh, arr[i][1])
	}
	return rooms
}

func main() {
	fmt.Println(meeting_rooms_two([][]int{{0, 30}, {5, 10}, {15, 20}}))
}
