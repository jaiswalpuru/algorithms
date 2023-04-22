package main

import "fmt"

func delete_tree_nodes(nodes int, parent []int, values []int) int {
	mp := make(map[int][]int)

	for i := 0; i < len(parent); i++ {
		if parent[i] == -1 {
			continue
		}
		mp[parent[i]] = append(mp[parent[i]], i)
	}

	sum_zero(mp, values, &nodes, 0)
	return nodes
}

func sum_zero(mp map[int][]int, values []int, nodes *int, root int) []int {
	res := []int{0, 0}
	for i := 0; i < len(mp[root]); i++ {
		r := sum_zero(mp, values, nodes, mp[root][i])
		res[0] += r[0]
		res[1] += r[1]
	}

	if res[0]+values[root] == 0 {
		*nodes -= (res[1] + 1)
		return []int{0, 0}
	}

	return []int{res[0] + values[root], res[1] + 1}
}

func main() {
	fmt.Println(delete_tree_nodes(7, []int{-1, 0, 0, 1, 2, 2, 2}, []int{1, -2, 4, 0, -2, -1, -1}))
}
