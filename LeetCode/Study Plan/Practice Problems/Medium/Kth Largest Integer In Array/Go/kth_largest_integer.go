package main

import (
	"container/heap"
	"fmt"
)

type STRING []string

func (mh STRING) Len() int { return len(mh) }

func (mh STRING) Less(i, j int) bool {
	if len(mh[i]) == len(mh[j]) {
		return mh[i] > mh[j]
	} else {
		return len(mh[i]) > len(mh[j])
	}
}

func (mh STRING) Swap(i, j int) { mh[i], mh[j] = mh[j], mh[i] }
func (mh *STRING) Pop() interface{} {
	var x interface{}
	x, *mh = (*mh)[len(*mh)-1], (*mh)[:len(*mh)-1]
	return x
}
func (mh *STRING) Push(i interface{}) { *mh = append(*mh, i.(string)) }

func kth_largest_integer(integers []string, k int) string {
	n := len(integers)
	str := make(STRING, n)
	for i := 0; i < n; i++ {
		str[i] = integers[i]
	}
	heap.Init(&str)

	var val string
	for k > 0 {
		val = heap.Pop(&str).(string)
		k--
	}
	return val
}

func main() {
	fmt.Println(kth_largest_integer([]string{"3", "6", "7", "10"}, 4))
}
