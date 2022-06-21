package main

import "fmt"
import "container/heap"

type N []int

func (mh N) Len() int              { return len(mh) }
func (mh N) Less(i, j int) bool    { return mh[i] < mh[j] }
func (mh N) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *N) Push(val interface{}) { *mh = append(*mh, val.(int)) }
func (mh *N) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func furthest_building_you_can_reach(arr []int, bricks, ladder int) int {
	n := len(arr)
	mh := &N{}
	heap.Init(mh)
	for i := 0; i < n-1; i++ {
		climb := arr[i+1] - arr[i]
		if climb <= 0 {
			continue
		}
		heap.Push(mh, climb)
		if mh.Len() <= ladder {
			continue
		}
		bricks -= heap.Pop(mh).(int)
		if bricks < 0 {
			return i
		}
	}
	return len(arr) - 1
}

func main() {
	fmt.Println(furthest_building_you_can_reach([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
}
