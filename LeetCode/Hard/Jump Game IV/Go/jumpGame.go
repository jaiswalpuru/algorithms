package main

import "fmt"

func minJumps(arr []int) int {
	size := len(arr)
	minJump := 0
	graph := make(map[int][]int)
	visited := make(map[int]bool)

	for i := 0; i < size; i++ {
		graph[arr[i]] = append(graph[arr[i]], i)
	}

	q := []int{0}
	for len(q) > 0 {
		qSize := len(q)
		for i := 0; i < qSize; i++ {
			curr := q[0]
			q = q[1:]

			if curr == size-1 {
				return minJump
			}

			for j := 0; j < len(graph[arr[curr]]); j++ {
				child := graph[arr[curr]][j]
				if !visited[child] {
					q = append(q, child)
					visited[child] = true
				}
			}

			delete(graph, arr[curr])

			if curr+1 < size && !visited[curr+1] {
				visited[curr+1] = true
				q = append(q, curr+1)
			}

			if curr-1 >= 0 && !visited[curr-1] {
				visited[curr-1] = true
				q = append(q, curr-1)
			}
		}
		minJump++
	}

	return -1
}

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}
