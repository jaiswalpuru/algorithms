package main

import (
	"container/heap"
	"fmt"
)

type P struct{ x, y, sum int }

type I []P

func (mh I) Len() int              { return len(mh) }
func (mh I) Less(i, j int) bool    { return mh[i].sum < mh[j].sum }
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(P)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func find_k_pairs_with_smallest_sum(a, b []int, k int) [][]int {
	m := I{}

	n := len(a)
	for i := 0; i < n; i++ {
		v := P{i, 0, a[i] + b[0]}
		heap.Push(&m, v)
	}

	res := [][]int{}
	for k > 0 && len(m) > 0 {
		val := heap.Pop(&m).(P)
		i, j := val.x, val.y
		res = append(res, []int{a[i], b[j]})
		if j+1 < len(b) {
			heap.Push(&m, P{i, j + 1, a[i] + b[j+1]})
		}
		k--
	}

	return res
}

func main() {
	fmt.Println(find_k_pairs_with_smallest_sum([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}
