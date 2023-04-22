package main

import "fmt"

var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func rotting_oranges(arr [][]int) int {
	n, m := len(arr), len(arr[0])
	fresh := 0
	q := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if arr[i][j] == 1 {
				fresh++
			}
		}
	}

	time := -1
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			curr := q[0]
			q = q[1:]
			for k := 0; k < 4; k++ {
				x, y := curr[0]+direction[k][0], curr[1]+direction[k][1]
				if is_safe(x, y, n, m, arr) {
					arr[x][y] = 2
					fresh--
					q = append(q, []int{x, y})
				}
			}
		}
		time++
	}

	if fresh > 0 {
		return -1
	}
	if time == -1 {
		return 0
	}
	return time
}

func is_safe(i, j, n, m int, arr [][]int) bool {
	if i < 0 || j < 0 || i >= n || j >= m || arr[i][j] != 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(rotting_oranges([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
}
