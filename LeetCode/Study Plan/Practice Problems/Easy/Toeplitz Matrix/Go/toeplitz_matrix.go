package main

import "fmt"

type Pair struct{ i, j int }

func toeplitz_matrix(arr [][]int) bool {
	n, m := len(arr), len(arr[0])
	q := []Pair{}
	for i := 0; i < n; i++ {
		q = append(q, Pair{i, 0})
	}
	for j := 1; j < m; j++ {
		q = append(q, Pair{0, j})
	}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.i+1 >= n || curr.j+1 >= m {
			continue
		}
		if arr[curr.i][curr.j] != arr[curr.i+1][curr.j+1] {
			return false
		}
		q = append(q, Pair{curr.i + 1, curr.j + 1})
	}
	return true
}

func main() {
	fmt.Println(toeplitz_matrix([][]int{
		{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2},
	}))
}
