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

//---------------------------------using quick select approach----------------
//considering the 0th index element as the pivot
func partition(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] >= arr[l+1] {
			arr[l+1], arr[l] = arr[l], arr[l+1]
			l++
		} else if arr[l] < arr[r] {
			r--
		} else {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	return l
}

func quick_select(arr []int, k int) int {
	p := partition(arr)
	if k == p {
		return arr[p]
	} else if k < p {
		return quick_select(arr[:p], k)
	} else {
		return quick_select(arr[p+1:], k-(p+1))
	}
}

func kth_largest_element_in_array(arr []int, k int) int {
	return quick_select(arr, len(arr)-k)
}

//---------------------------------using quick select approach----------------

func main() {
	fmt.Println(kth_largest_element_in_array([]int{3, 2, 1, 5, 6, 4}, 2))
}
