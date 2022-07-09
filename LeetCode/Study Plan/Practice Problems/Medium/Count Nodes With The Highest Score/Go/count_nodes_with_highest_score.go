package main

import "fmt"

func make_tree(arr []int) map[int][]int {
	tree := make(map[int][]int)

	n := len(arr)
	for i := 1; i < n; i++ {
		tree[arr[i]] = append(tree[arr[i]], i)
	}
	return tree
}

func count_the_highest_score(parents []int) int {
	tree := make_tree(parents)
	n, cnt := 0, 0
	dfs(parents, tree, 0, &n, &cnt)
	return cnt
}

func dfs(parents []int, tree map[int][]int, curr int, n, cnt *int) int {
	c := []int{0, 0}
	for i := 0; i < len(tree[curr]); i++ {
		c[i] = dfs(parents, tree, tree[curr][i], n, cnt)
	}

	res := len(parents) - 1 - c[0] - c[1]
	if res == 0 {
		res = 1
	}
	for _, v := range c {
		if v == 0 {
			v = 1
		}
		res *= v
	}
	if res > *n {
		*n = res
		*cnt = 1
	} else if res == *n {
		*cnt++
	}
	return c[0] + c[1] + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(count_the_highest_score([]int{-1, 2, 0, 2, 0}))
}
