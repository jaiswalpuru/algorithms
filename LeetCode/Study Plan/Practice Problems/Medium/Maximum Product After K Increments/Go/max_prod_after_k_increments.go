package main

import (
	"container/heap"
	"fmt"
)

var mod = 1e9 + 7

type I []int

func (mh I) Len() int              { return len(mh) }
func (mh I) Less(i, j int) bool    { return mh[i] < mh[j] }
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(int)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func maximum_product_after_k_increments(arr []int, k int) int {
	var a I
	a = arr
	heap.Init(&a)

	for k > 0 {
		val := heap.Pop(&a).(int)
		val += 1
		heap.Push(&a, val)
		k--
	}
	n := len(a)
	res := 1
	for i := 0; i < n; i++ {
        res = (res*a[i])%int(mod)
	}
	return res
}

func main() {
	fmt.Println(maximum_product_after_k_increments([]int{0, 4}, 5))
}
