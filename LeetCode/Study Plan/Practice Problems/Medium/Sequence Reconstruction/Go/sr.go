package main

import "fmt"

func sequence_reconstruct(num []int, seq [][]int) bool {
	graph := make(map[int][]int)
	indegree := make(map[int]int)
	n := len(seq)

	for i := 0; i < n; i++ {
		for _, v := range seq[i] {
			indegree[v] = 0
		}
	}

	if len(indegree) != len(num) {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 1; j < len(seq[i]); j++ {
			prev := seq[i][j-1]
			curr := seq[i][j]
			graph[prev] = append(graph[prev], curr)
			indegree[curr]++
		}
	}

	var q []int
	for k, v := range indegree {
		if v == 0 {
			q = append(q, k)
		}
	}

	res := []int{}
	for len(q) > 0 {
		if len(q) > 1 {
			return false
		}
		curr := q[0]
		if curr != num[len(res)] {
			return false
		}
		res = append(res, curr)
		q = q[1:]

		for i := 0; i < len(graph[curr]); i++ {
			indegree[graph[curr][i]]--
			if indegree[graph[curr][i]] == 0 {
				q = append(q, graph[curr][i])
			}
		}
	}

	return len(res) == len(num)
}

func main() {
	fmt.Println(sequence_reconstruct([]int{1, 2, 3}, [][]int{
		{1, 2}, {1, 3},
	}))
}
