package main

import "fmt"

func minimum_hamming_distance_after_swap_operations(source, target []int, allowed_swaps [][]int) int {
	parent := make([]int, len(source))
	size := make([]int, len(source))
	for i := 0; i < len(source); i++ {
		parent[i] = i
		size[i] = 1
	}
	for i := 0; i < len(allowed_swaps); i++ {
		union(allowed_swaps[i][0], allowed_swaps[i][1], &parent, &size)
	}
	group := make(map[int][]int)
	for i := 0; i < len(parent); i++ {
		g := find(i, &parent)
		group[g] = append(group[g], i)
	}
	cnt := 0
	for _, curr := range group {
		if len(curr) != 0 {
			src_cnt := make(map[int]int)
			tgt_cnt := make(map[int]int)
			for j := 0; j < len(curr); j++ {
				src_cnt[source[curr[j]]]++
				tgt_cnt[target[curr[j]]]++
			}
			for k, v := range tgt_cnt {
				if v >= src_cnt[k] {
					cnt += v - src_cnt[k]
				}
			}
		}
	}
	return cnt
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] >= (*size)[y] {
		(*size)[x] += (*size)[y]
		(*parent)[y] = x
	} else {
		(*size)[y] += (*size)[x]
		(*parent)[x] = y
	}
}

func find(x int, parent *[]int) int {
	if x != (*parent)[x] {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
}

func main() {
	fmt.Println(minimum_hamming_distance_after_swap_operations([]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{
		{0, 1}, {2, 3},
	}))
}
