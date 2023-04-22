package main

import "fmt"

func valid(n int, left, right []int) bool {
	root := make([]bool, n)
	for i := 0; i < n; i++ {
		root[i] = true
	}

	for i := 0; i < n; i++ {
		if left[i] != -1 {
			root[left[i]] = false
		}
		if right[i] != -1 {
			root[right[i]] = false
		}
	}

	root_cnt := 0
	_root := -1
	for i := 0; i < n; i++ {
		if root[i] {
			root_cnt++
			_root = i
		}
	}

	if root_cnt != 1 {
		return false
	}

	visited := make(map[int]bool)
	q := []int{_root}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if visited[curr] == true {
			return false
		}
		visited[curr] = true

		if left[curr] != -1 {
			q = append(q, left[curr])
		}

		if right[curr] != -1 {
			q = append(q, right[curr])
		}
	}
	return len(visited) == n
}

func main() {
	fmt.Println(valid(6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}))
}
