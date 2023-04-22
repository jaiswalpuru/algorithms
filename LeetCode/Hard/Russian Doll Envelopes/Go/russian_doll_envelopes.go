package main

import "fmt"
import "sort"

func max(a, b int) int {
	if a>b{
		return a
	}
	return b
}

type I [][]int
func (a I) Len()int { return len(a) }
func (a I) Swap(i, j int) { a[i],a[j] = a[j],a[i] }
func (a I) Less(i, j int) bool { 
	if a[i][0] == a[j][0] {
		return a[i][1] > a[j][1]
	}else {
		return a[i][0] < a[j][0]
	}
}

//-------------------------Brute force-----------------------------
func russian_doll_envelopes_brute(env [][]int) int {
	var arr I
    arr = env
    sort.Sort(arr)
    return _recur(env, 0, []int{-1,-1})
}

func _recur(arr [][]int, ind int, prev []int) int {
	if ind == len(arr) {
		return 0
	}

	exclude := _recur(arr, ind+1, prev)
	include := 0
	if arr[ind][0] > prev[0] && arr[ind][1] > prev[1] {
		include = 1 + _recur(arr, ind+1, arr[ind])
	}
	return max(exclude, include)
}
//-------------------------Brute force-----------------------------

//----------------------------binary search approach------------------------------
func russian_dolls_binary_search(env [][]int) int {
	sort.Sort(I(env))
	ht := make([]int, len(env))
	for i:=0;i<len(env);i++{
		ht[i] = env[i][1]
	}
	return lis(ht)
}

func lis(ht []int) int {
	dp := make([]int, len(ht))
	l := 0
	for i:=0;i<len(ht);i++{
		pos := sort.SearchInts(dp[:l], ht[i])
		dp[pos] = ht[i]
		if pos == l {
			l++
		}
	}
	return l
}
//----------------------------binary search approach------------------------------

func main() {
	fmt.Println(russian_dolls_binary_search([][]int{
		{5,4},{6,4},{6,7},{2,3},
	}))
}