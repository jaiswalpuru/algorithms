package main

import (
	"fmt"
)

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum_domino_rotation(tops, bottom []int) int {
	max_top := 0
	max_top_cnt := 0
	max_bottom := 0
	max_bottom_cnt := 0
	hash_map_top := make(map[int]int)
	hash_map_bottom := make(map[int]int)

	n := len(tops)

	for i:=0 ; i<n; i++ {
		
		hash_map_top[tops[i]]++
		if max_top_cnt < hash_map_top[tops[i]] {
			max_top_cnt = hash_map_top[tops[i]]
			max_top = tops[i]
		}

		hash_map_bottom[bottom[i]]++
		if max_bottom_cnt < hash_map_bottom[bottom[i]] {
			max_bottom_cnt = hash_map_bottom[bottom[i]]
			max_bottom = bottom[i]
		}
	}

	min_swap_top := 0
	for i:=0; i<n; i++ {
		if tops[i] != max_top {
			if bottom[i] == max_top {
				min_swap_top++
			}else{
				min_swap_top = -1
				break
			}
		}
	}

	min_swap_bottom := 0
	for i:=0;i<n;i++{
		if bottom[i] != max_bottom { 
			if tops[i] == max_bottom {
				min_swap_bottom++
			}else{
				min_swap_bottom = -1
				break
			}
		}
	}

	if min_swap_bottom == -1 && min_swap_top == -1 {
		return -1
	}else{
		if min_swap_bottom == -1 {
			return min_swap_top
		}else if min_swap_top == -1 {
			return min_swap_bottom
		}else {
			return min(min_swap_bottom, min_swap_top)
		}
	}
}

func min(a,b int) int {
	if a>b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimum_domino_rotation([]int{2,1,2,4,2,2}, []int{5,2,6,2,3,2}))
}