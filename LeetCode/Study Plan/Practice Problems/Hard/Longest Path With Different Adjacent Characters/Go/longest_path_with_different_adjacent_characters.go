package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(node int, hash_map map[int][]int, res *int, s string) int {
	two_max := []int{0, 0}
	for i := 0; i < len(hash_map[node]); i++ {
		curr := dfs(hash_map[node][i], hash_map, res, s)
		if s[node] != s[hash_map[node][i]] {
			if curr > two_max[0] {
				two_max[1] = two_max[0]
				two_max[0] = curr
			} else if curr > two_max[1] {
				two_max[1] = curr
			}
		}
	}
	*res = max(*res, two_max[0]+two_max[1]+1)
	return two_max[0] + 1
}

func longest_path_with_different_adjacent_characters(parent []int, s string) int {
	hash_map := make(map[int][]int)
	for i := 0; i < len(parent); i++ {
		if parent[i] >= 0 {
			hash_map[parent[i]] = append(hash_map[parent[i]], i)
		}
	}
	res := 0
	dfs(0, hash_map, &res, s)
	return res
}

func main() {
	fmt.Println(longest_path_with_different_adjacent_characters([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
}
