package main

import (
	"fmt"
	"strconv"
)

func count_artifacts(n int, artifacts [][]int, dig [][]int) int {
	hash_map := make(map[string]bool)
	m := len(dig)
	for i := 0; i < m; i++ {
		hash_map[to_string(dig[i][0], dig[i][1])] = true
	}

	cnt := 0
	for i := 0; i < len(artifacts); i++ {
		r1, c1, r2, c2 := artifacts[i][0], artifacts[i][1], artifacts[i][2], artifacts[i][3]
		f := true
		for i := r1; i <= r2; i++ {
			for j := c1; j <= c2; j++ {
				if !hash_map[to_string(i, j)] {
					f = false
					break
				}
			}
		}
		if f {
			cnt++
		}
	}
	return cnt
}

func to_string(a, b int) string { return strconv.Itoa(a) + "->" + strconv.Itoa(b) }

func main() {
	fmt.Println(count_artifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}}))
}
