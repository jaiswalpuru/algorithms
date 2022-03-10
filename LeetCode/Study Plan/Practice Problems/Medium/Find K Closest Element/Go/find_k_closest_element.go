package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type P struct{ v, d int }

type I []P

func (mh I) Len() int { return len(mh) }
func (mh I) Less(i, j int) bool {
	if mh[i].d == mh[j].d {
		return mh[i].v < mh[j].v
	}
	return mh[i].d < mh[j].d
}
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(P)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func find_closest_k_ele(arr []int, k, x int) []int {
	mh := &I{}
	n := len(arr)
	for i := 0; i < n; i++ {
		heap.Push(mh, P{arr[i], abs(arr[i] - x)})
	}
	res := []int{}
	for mh.Len() > 0 && k > 0 {
		res = append(res, heap.Pop(mh).(P).v)
		k--
	}
	sort.Ints(res)
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	fmt.Println(find_closest_k_ele([]int{1, 2, 3, 4, 5}, 4, 3))
}
