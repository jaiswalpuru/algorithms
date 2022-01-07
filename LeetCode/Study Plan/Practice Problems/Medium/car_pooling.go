/*
There is a car with capacity empty seats. The vehicle only drives east (i.e., it cannot turn around and drive west).

You are given the integer capacity and an array trips where trip[i] = [numPassengersi, fromi, toi] indicates that the
ith trip has numPassengersi passengers and the locations to pick them up and drop them off are fromi and toi respectively.
The locations are given as the number of kilometers due east from the car's initial location.

Return true if it is possible to pick up and drop off all passengers for all the given trips, or false otherwise.
*/

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap [][]int

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh MinHeap) Less(i, j int) bool { return mh[i][2] < mh[j][2] }

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.([]int))
}

func (mh *MinHeap) Pop() interface{} {
	var x interface{}
	x, *mh = (*mh)[len(*mh)-1], (*mh)[:len(*mh)-1]
	return x
}

func car_pooling(arr [][]int, cap int) bool {
	n := len(arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})

	mh := MinHeap{}
	heap.Init(&mh)

	for i := 0; i < n; i++ {
		heap.Push(&mh, arr[i])
		for mh[0][2] <= arr[i][1] {
			cap += mh[0][0]
			heap.Remove(&mh, 0)
		}
		cap -= arr[i][0]
		if cap < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(car_pooling([][]int{{9, 3, 6}, {8, 1, 7}, {6, 6, 8}, {8, 4, 9}, {4, 2, 9}}, 28))
}
