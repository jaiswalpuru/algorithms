package main

import (
	"container/heap"
	"fmt"
)

type P struct {
	val int
	cnt int
}

type I []P

func (mh I) Len() int              { return len(mh) }
func (mh I) Less(i, j int) bool    { return mh[i].cnt > mh[j].cnt }
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(P)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func top_k_frequent_element(arr []int, k int) []int {
	res := []int{}
	visited := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		visited[arr[i]]++
	}

	mh := &I{}
	for key, v := range visited {
		heap.Push(mh, P{key, v})
	}

	for mh.Len() > 0 && k > 0 {
		res = append(res, heap.Pop(mh).(P).val)
		k--
	}
	return res
}

func main() {
	fmt.Println(top_k_frequent_element([]int{1, 1, 1, 2, 2, 3}, 2))
}
