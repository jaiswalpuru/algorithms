package main

import (
	"fmt"
	"math"
	"strconv"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func to_string(arr []int, i int) string {
	return strconv.Itoa(arr[0]) + strconv.Itoa(arr[1]) + strconv.Itoa(arr[2]) + strconv.Itoa(i)
}

func euclid_dis(x1, y1, x2, y2, radius int) bool {
	return math.Sqrt(float64(abs(x2-x1)*abs(x2-x1))+float64(abs(y2-y1)*abs(y2-y1))) <= float64(radius)
}

func detonate(bombs [][]int) int {
	n := len(bombs)
	g := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				if euclid_dis(bombs[i][0], bombs[i][1], bombs[j][0], bombs[j][1], bombs[i][2]) {
					g[i] = append(g[i], j)
				}
			}
		}
	}
	ans := -1
	for i := 0; i < n; i++ {
		res := 0
		visited := make(map[int]struct{})
		dfs(g, i, &res, &visited)
		ans = max(ans, res)

	}
	return ans
}

func dfs(g map[int][]int, curr int, res *int, visited *map[int]struct{}) {
	(*visited)[curr] = struct{}{}
	*res++
	for i := 0; i < len(g[curr]); i++ {
		if _, ok := (*visited)[g[curr][i]]; !ok {
			dfs(g, g[curr][i], res, visited)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(detonate([][]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 4, 2},
		{4, 5, 3},
		{5, 6, 4},
	}))
}
