package main

import (
	"container/heap"
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------brute force------------------------
func jump_game(arr []int, k int) int {
	res := math.MinInt64
	_recur(arr, k, 0, arr[0], &res)
	return res

}

func _recur(arr []int, k int, ind int, sum int, res *int) {
	if ind == len(arr)-1 {
		*res = max(*res, sum)
		return
	}

	for i := 1; i <= k && (ind+i) < len(arr); i++ {
		_recur(arr, k, ind+i, sum+arr[ind+i], res)
	}
}

//------------------brute force------------------------

//------------------memoization still TLE------------------------
func jump_game_memo(arr []int, k int) int {
	memo := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		memo[i] = -1
	}
	return _memo(arr, k, 0, &memo)
}

func _memo(arr []int, k, ind int, memo *[]int) int {
	if (*memo)[ind] != -1 {
		return (*memo)[ind]
	}
	if ind == len(arr)-1 {
		(*memo)[ind] = arr[ind]
		return (*memo)[ind]
	}

	max_val := math.MinInt64
	for i := 1; i <= k && (ind+i) < len(arr); i++ {
		max_val = max(max_val, _memo(arr, k, ind+i, memo))
	}
	(*memo)[ind] = max_val + arr[ind]
	return (*memo)[ind]
}

//------------------memoization still TLE------------------------

//------------------efficient approach---------------------------
type Pair struct{ score, ind int }

type I []Pair

func (mh I) Len() int              { return len(mh) }
func (mh I) Less(i, j int) bool    { return mh[i].score > mh[j].score }
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(Pair)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func jump_game_eff(arr []int, k int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	mh := &I{}
	n := len(arr)
	heap.Push(mh, Pair{arr[0], 0})
	for i := 1; i < n; i++ {
		for (*mh)[0].ind < i-k {
			heap.Pop(mh)
		}
		dp[i] = arr[i] + dp[(*mh)[0].ind]
		heap.Push(mh, Pair{dp[i], i})
	}
	return dp[n-1]
}

//------------------efficient approach---------------------------

func main() {
	fmt.Println(jump_game_eff([]int{1, -1, -2, 4, -7, 3}, 2))
}
