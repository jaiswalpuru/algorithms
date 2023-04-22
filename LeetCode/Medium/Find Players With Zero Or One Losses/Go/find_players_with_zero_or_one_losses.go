package main

import (
	"fmt"
	"sort"
)

func find_winners(matches [][]int) [][]int {
	n := len(matches)
	hash_won := make(map[int]int)
	hash_lost := make(map[int]int)

	for i := 0; i < n; i++ {
		hash_won[matches[i][0]]++
		hash_lost[matches[i][1]]++
	}

	won := []int{}
	lost := []int{}

	for k := range hash_won {
		if _, ok := hash_lost[k]; !ok {
			won = append(won, k)
		}
	}

	for k, v := range hash_lost {
		if v == 1 {
			lost = append(lost, k)
		}
	}

	sort.Ints(won)
	sort.Ints(lost)
	return [][]int{won, lost}
}

func main() {
	fmt.Println(find_winners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
}
