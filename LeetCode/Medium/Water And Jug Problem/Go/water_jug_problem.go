package main

import (
	"fmt"
)

type Pair struct {
	j1, j2 int
}

func water_jug_problem(j1, j2, target int) bool {
	if j1+j2 < target {
		return false
	}
	visited := make(map[Pair]bool)
	q := []Pair{{0, 0}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.j1+curr.j2 == target {
			return true
		}
		temp := []Pair{}
		temp = append(temp, Pair{curr.j1, 0})
		temp = append(temp, Pair{0, curr.j2})
		temp = append(temp, Pair{j1, curr.j2})
		temp = append(temp, Pair{curr.j1, j2})
		if curr.j2 < j1-curr.j1 {
			temp = append(temp, Pair{min(curr.j1, curr.j1+curr.j2), 0})
		} else {
			temp = append(temp, Pair{min(curr.j1, curr.j1+curr.j2), curr.j2 - (j1 - curr.j1)})
		}
		if curr.j1 < j2-curr.j2 {
			temp = append(temp, Pair{0, min(j2, curr.j1+curr.j2)})
		} else {
			temp = append(temp, Pair{curr.j1 - (j2 - curr.j2), min(j2, curr.j1+curr.j2)})
		}
		for i := 0; i < len(temp); i++ {
			if visited[temp[i]] {
				continue
			}
			q = append(q, temp[i])
			visited[temp[i]] = true
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(water_jug_problem(3, 5, 4))
}
