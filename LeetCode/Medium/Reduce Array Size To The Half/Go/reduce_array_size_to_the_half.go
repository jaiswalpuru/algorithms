package main

import (
	"container/heap"
	"fmt"
)

type Pair struct{ val, cnt int }

type MH []Pair

func (m MH) Len() int              { return len(m) }
func (m MH) Less(i, j int) bool    { return m[i].cnt > m[j].cnt }
func (m MH) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *MH) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *MH) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func reduce_array_size_to_the_half(arr []int) int {
	mp := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}
	mh := &MH{}
	for k, v := range mp {
		heap.Push(mh, Pair{k, v})
	}
	res := 0
	for mh.Len() > 0 {
		c := heap.Pop(mh).(Pair)
		n -= c.cnt
		res++
		if n <= len(arr)/2 {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(reduce_array_size_to_the_half([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}
