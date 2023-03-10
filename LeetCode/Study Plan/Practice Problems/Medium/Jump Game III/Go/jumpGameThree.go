package main

import "fmt"

func canReach(arr []int, start int) bool {
	q := []int{start}
	size := len(arr)
	visited := make(map[int]bool)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if arr[curr] == 0 {
			return true
		}

		if curr+arr[curr] < size {
			if !visited[curr+arr[curr]] {
				visited[curr+arr[curr]] = true
				q = append(q, curr+arr[curr])
			}
		}

		if curr-arr[curr] >= 0 {
			if !visited[curr-arr[curr]] {
				visited[curr-arr[curr]] = true
				q = append(q, curr-arr[curr])
			}
		}
	}

	return false
}

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
}
