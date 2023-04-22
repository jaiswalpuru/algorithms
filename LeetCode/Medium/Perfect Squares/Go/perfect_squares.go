package main

import (
	"fmt"
	"math"
)

//---------------------recursive(TLE)---------------------------
func perfect_squares(n int) int {
	mp := make(map[int]int)
	sq := make([]int, int(math.Sqrt(float64(n))+1))
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		mp[i*i] = i * i
		sq[i] = i * i
	}
	return min_num_sq(mp, sq, n)
}

func min_num_sq(mp map[int]int, sq []int, n int) int {
	if mp[n] != 0 {
		return 1
	}
	min_num := math.MaxInt64
	for i := 1; i < len(sq); i++ {
		if n < sq[i] {
			break
		}
		val := min_num_sq(mp, sq, n-sq[i]) + 1
		min_num = min(min_num, val)
	}
	return min_num
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//---------------------recursive(TLE)---------------------------

//---------------------memo---------------------------
func perfect_squares_memo(n int) int {
	memo := make(map[int]int)
	mp := make(map[int]int)
	sq := make([]int, int(math.Sqrt(float64(n))+1))
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		mp[i*i] = i * i
		sq[i] = i * i
	}
	return min_num_memo(&memo, sq, mp, n)
}

func min_num_memo(memo *map[int]int, sq []int, mp map[int]int, n int) int {
	if (*memo)[n] != 0 {
		return (*memo)[n]
	}
	if mp[n] != 0 {
		return 1
	}
	min_cnt := math.MaxInt64
	for i := 1; i < len(sq); i++ {
		if n < sq[i] {
			break
		}
		val := min_num_memo(memo, sq, mp, n-sq[i]) + 1
		min_cnt = min(min_cnt, val)
		(*memo)[n] = min_cnt
	}
	return min_cnt
}

//---------------------memo---------------------------

//------------------------a better approach-----------------------
func perfect_squares_eff(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt64
	}
	max_ind := int(math.Sqrt(float64(n))) + 1
	sq := make([]int, max_ind)
	for i := 1; i < max_ind; i++ {
		sq[i] = i * i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j < max_ind; j++ {
			if i < sq[j] {
				break
			}
			dp[i] = min(dp[i], dp[i-sq[j]]+1)
		}
	}
	return dp[n]
}

//------------------------a better approach-----------------------

func main() {
	fmt.Println(perfect_squares_eff(12))
}
