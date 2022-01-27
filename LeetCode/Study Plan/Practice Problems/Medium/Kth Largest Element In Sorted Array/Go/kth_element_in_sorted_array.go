package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//using heaps in go for using inbuilt max heaps we need to first define some functions which satisfies the interface
type INT []int

func (mh INT) Len() int           { return len(mh) }
func (mh INT) Less(i, j int) bool { return mh[i] > mh[j] }
func (mh INT) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *INT) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func (mh *INT) Push(val interface{}) { *mh = append(*mh, val.(int)) }

func kth_largest_element(arr []int, k int) int {
	n := len(arr)
	t := make(INT, n)
	for i := 0; i < n; i++ {
		t[i] = arr[i]
	}

	heap.Init(&t)
	var val int
	for k > 0 {
		val = heap.Pop(&t).(int)
		k--
	}
	return val
}

//method 1 sort it and return
func largest_kth_ele(arr []int, k int) int {
	sort.Ints(arr)
	return arr[len(arr)-k]
}

func main() {
	fmt.Println(kth_largest_element([]int{3, 2, 1, 5, 6, 4}, 2))
}
