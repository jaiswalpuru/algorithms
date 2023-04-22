package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	size := len(rooms)
	visited := make(map[int]bool)
	q := []int{0}

	for len(q) > 0 {
		qSize := len(q)
		for i := 0; i < qSize; i++ {
			key := q[0]
			q = q[1:]
			visited[key] = true
			for j := 0; j < len(rooms[key]); j++ {
				if !visited[rooms[key][j]] {
					q = append(q, rooms[key][j])
				}
			}
		}
	}

	for i := 0; i < size; i++ {
		if !visited[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{
		{1, 3}, {3, 0, 1}, {2}, {0},
	}))
}
