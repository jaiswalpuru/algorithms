package main

import (
	"container/heap"
	"fmt"
	"math"
)

//max heap
type P []int

func (mh P) Len() int              { return len(mh) }
func (mh P) Less(i, j int) bool    { return mh[i] > mh[j] }
func (mh P) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *P) Push(val interface{}) { *mh = append(*mh, val.(int)) }
func (mh *P) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimize_deviation(arr []int) int {

	maxh := &P{}
	n := len(arr)
	min_eve := math.MaxInt64
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			min_eve = min(min_eve, arr[i])
			heap.Push(maxh, arr[i])
		} else {
			min_eve = min(min_eve, arr[i]*2)
			heap.Push(maxh, arr[i]*2)
		}
	}

	min_deviation := math.MaxInt64

	for (*maxh).Len() > 0 {
		val := heap.Pop(maxh).(int)
		min_deviation = min(min_deviation, val-min_eve)
		if val%2 == 0 {
			heap.Push(maxh, val/2)
			min_eve = min(min_eve, val/2)
		} else {
			break
		}
	}
	return min_deviation
}

func main() {
	fmt.Println(minimize_deviation([]int{1, 2, 3, 4}))
}
