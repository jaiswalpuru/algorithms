package main

import (
	"fmt"
	"strconv"
)

func water_jug_problem(j1, j2, target int) bool {
	if j1+j2 < target {
		return false
	}
	visited := make(map[string]bool)
	q := []Pair{{0, 0}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.j1+curr.j2 == target {
			return true
		}
		temp := []Pair{}
		temp = append(temp, Pair{j1, curr.j2}) //fill jar 1
		temp = append(temp, Pair{curr.j1, j2}) //fill jar 2
		temp = append(temp, Pair{0, curr.j2})  //empty jar 1
		temp = append(temp, Pair{curr.j1, 0})  //empty jar 2
		if curr.j2 < j1-curr.j1 {
			temp = append(temp, Pair{min(j1, curr.j1+curr.j2), 0})
		} else {
			temp = append(temp, Pair{min(j1, curr.j1+curr.j2), curr.j2 - (j1 - curr.j1)})
		}
		if curr.j1 < j2-curr.j2 {
			temp = append(temp, Pair{0, min(j2, curr.j1+curr.j2)})
		} else {
			temp = append(temp, Pair{curr.j1 - (j2 - curr.j2), min(j2, curr.j1+curr.j2)})
		}
		for i := 0; i < len(temp); i++ {
			if visited[tos(temp[i].j1, temp[i].j2)] {
				continue
			}
			q = append(q, Pair{temp[i].j1, temp[i].j2})
			visited[tos(temp[i].j1, temp[i].j2)] = true
		}
	}
	return false

}

func tos(a, b int) string {
	return strconv.Itoa(a) + ":" + strconv.Itoa(b)
}

type Pair struct {
	j1, j2 int
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
