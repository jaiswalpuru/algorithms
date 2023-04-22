package main

import "fmt"
import "container/heap"

type I []int

func (mh I) Len() int              { return len(mh) }
func (mh I) Less(i, j int) bool    { return mh[i] > mh[j] }
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(int)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	(*mh) = (*mh)[:len(*mh)-1]
	return val
}

//--------------------------TLE--------------------------------- refer to LC solutions
func construct_target_array_with_multiple_sum(arr []int) bool {
	if len(arr) == 1 {
		return arr[0] == 1
	}
	mh := &I{}
	heap.Init(mh)
	tot_sum := 0
	for _, num := range arr {
		tot_sum += num
		heap.Push(mh, num)
	}

	for (*mh)[0] > 1 {
		val := heap.Pop(mh).(int)
		diff := val - (tot_sum - val)
		if diff < 1 {
			return false
		}
		heap.Push(mh, diff)
		tot_sum = tot_sum - val + diff
	}

	return true

}

//--------------------------TLE---------------------------------

//--------------------------much better approach basically optimized approach to the above one---------------------------------
func construct_target_array_with_multiple_sum_eff(arr []int) bool {
	if len(arr) == 1 {
		return arr[0] == 1
	}
	mh := &I{}
	heap.Init(mh)
	tot_sum := 0
	for _, num := range arr {
		tot_sum += num
		heap.Push(mh, num)
	}

	for (*mh)[0] > 1 {
		max_val := heap.Pop(mh).(int)
		rest := tot_sum - max_val
		if rest == 1 {
			return true
		}
		x := max_val % rest
		if x == 0 || x == max_val {
			return false
		}
		heap.Push(mh, x)
		tot_sum = tot_sum - max_val + x
	}

	return true

}

//--------------------------much better approach basically optimized approach to the above one---------------------------------

func main() {
	fmt.Println(construct_target_array_with_multiple_sum_eff([]int{9, 3, 5}))
}
