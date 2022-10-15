package main

import (
	"fmt"
	"math"
)

//----------------------brute force-----------------------------------
func string_compression(s string, k int) int {
	return recur(s, 0, 'a'-1, 0, k)
}

func recur(s string, ind int, last_char byte, last_char_cnt int, k int) int {
	if k < 0 {
		return math.MaxInt
	}

	if ind == len(s) {
		return 0
	}

	delete_char := recur(s, ind+1, last_char, last_char_cnt, k-1)
	keep := 0
	if s[ind] == last_char {
		keep = recur(s, ind+1, last_char, last_char_cnt+1, k)
		if last_char_cnt == 1 || last_char_cnt == 9 || last_char_cnt == 99 {
			keep++
		}
	} else {
		keep = recur(s, ind+1, s[ind], 1, k) + 1
	}

	return min(keep, delete_char)
}

//----------------------brute force-----------------------------------

//----------------------memoize-----------------------------------
type S struct {
	ind           int
	last_char     byte
	last_char_cnt int
	k             int
}

func string_compression_memo(s string, k int) int {
	memo := make(map[S]int)
	return _memo(s, 0, 'a'-1, 0, k, &memo)
}

func _memo(s string, ind int, last_char byte, last_char_cnt int, k int, memo *map[S]int) int {
	state := S{ind, last_char, last_char_cnt, k}
	if v, ok := (*memo)[state]; ok {
		return v
	}
	if k < 0 {
		return math.MaxInt
	}
	if ind == len(s) {
		return 0
	}
	delete_char := _memo(s, ind+1, last_char, last_char_cnt, k-1, memo)
	keep := 0
	if s[ind] == last_char {
		keep = _memo(s, ind+1, last_char, last_char_cnt+1, k, memo)
		if last_char_cnt == 1 || last_char_cnt == 9 || last_char_cnt == 99 {
			keep++
		}
	} else {
		keep = _memo(s, ind+1, s[ind], 1, k, memo) + 1
	}
	(*memo)[state] = min(keep, delete_char)
	return (*memo)[state]
}

//----------------------memoize-----------------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(string_compression_memo("aaabcccd", 2))
}
