package main

import (
	"fmt"
	"math"
	"sort"
)

//------------------------------recursion(dont even bother trying this TLE)------------------------------------
func delete_and_earn_recursion(arr []int) int {
	count := make(map[int]int)
	n := len(arr)
	a := []int{}
	for i := 0; i < n; i++ {
		if _, ok := count[arr[i]]; !ok {
			a = append(a, arr[i])
		}
		count[arr[i]]++
	}
	sort.Ints(a)
	sum := math.MinInt64
	visit := make(map[int]bool)
	_delete_and_earn(a, count, 0, &sum, 0, &visit)
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _delete_and_earn(arr []int, count map[int]int, ind int, sum *int, temp int, visit *map[int]bool) {
	if ind == len(arr) {
		*sum = max(*sum, temp)
		return
	}

	if !(*visit)[arr[ind]] {
		t := temp + arr[ind]*(count)[arr[ind]]
		(*visit)[arr[ind]+1] = true
		(*visit)[arr[ind]-1] = true
		(*visit)[arr[ind]] = true
		_delete_and_earn(arr, count, ind+1, sum, t, visit)
		t -= (arr[ind] * (count)[arr[ind]])
		(*visit)[arr[ind]+1] = false
		(*visit)[arr[ind]-1] = false
		(*visit)[arr[ind]] = false
	}
	_delete_and_earn(arr, count, ind+1, sum, temp, visit)
}

//------------------------------------------------------------------------------------------------

//----------------------------------------memoization----------------------------------------------
func delete_and_earn_memo(arr []int) int {
	count := make(map[int]int)
	n := len(arr)
	max_num := 0
	for i := 0; i < n; i++ {
		max_num = max(max_num, arr[i])
		count[arr[i]] += arr[i]
	}
	cache := make(map[int]int)
	return _delete_and_earn_memo(max_num, &count, &cache)
}

func _delete_and_earn_memo(num int, count *map[int]int, cache *map[int]int) int {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return (*count)[num]
	}

	if _, ok := (*cache)[num]; ok {
		return (*cache)[num]
	}

	gain := (*count)[num]
	(*cache)[num] = max(_delete_and_earn_memo(num-1, count, cache), _delete_and_earn_memo(num-2, count, cache)+gain)
	return (*cache)[num]
}

//-----------------------------------------------------------------------------------------------

//------------------------------------------dp---------------------------------------------------
func delete_and_earn_dp(arr []int) int {
	count := make(map[int]int)
	max_num := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		count[arr[i]] += arr[i]
		max_num = max(max_num, arr[i])
	}

	cache := make([]int, max_num+1)
	cache[1] = count[1]

	for i := 2; i < len(cache); i++ {
		gain := count[i]
		cache[i] = max(cache[i-1], cache[i-2]+gain)
	}
	return cache[max_num]
}

//-----------------------------------------------------------------------------------------------

func main() {
	fmt.Println(delete_and_earn_dp([]int{3, 4, 2}))
}
