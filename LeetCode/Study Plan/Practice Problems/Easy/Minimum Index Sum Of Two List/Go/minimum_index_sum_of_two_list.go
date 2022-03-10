package main

import (
	"fmt"
	"math"
)

func minimum_index_sum_of_two_list(l1, l2 []string) []string {
	mp := make(map[string]int)
	n, m := len(l1), len(l2)
	for i := 0; i < n; i++ {
		mp[l1[i]] = i
	}
	res := make(map[int][]string)
	min_sum := math.MaxInt64
	for j := 0; j < m; j++ {
		if _, ok := mp[l2[j]]; ok {
			if min_sum > j+mp[l2[j]] {
				min_sum = j + mp[l2[j]]
			}
			res[j+mp[l2[j]]] = append(res[j+mp[l2[j]]], l2[j])
		}
	}
	return res[min_sum]
}

func main() {
	fmt.Println(minimum_index_sum_of_two_list([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}
