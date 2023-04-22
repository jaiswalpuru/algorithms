package main

import "fmt"
import "container/heap"

type P struct{ x, y, cost int }

var dir = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func is_safe(x, y, n, m int) bool { return x >= 0 && y >= 0 && x < n && y < m }

type PQ []P

func (m PQ) Len() int              { return len(m) }
func (m PQ) Less(i, j int) bool    { return m[i].cost < m[j].cost }
func (m PQ) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *PQ) Push(val interface{}) { *m = append(*m, val.(P)) }
func (m *PQ) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func minimum_obstacle_removal_to_reach_corner(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = 1e5 + 1
		}
	}

	mh := PQ{P{0, 0, 0}}
	heap.Init(&mh)

	for mh.Len() > 0 {
		curr := heap.Pop(&mh).(P)
		if curr.x == n-1 && curr.y == m-1 {
			return curr.cost
		}

		for i := 0; i < 4; i++ {
			x, y := curr.x+dir[i][0], curr.y+dir[i][1]
			if is_safe(x, y, n, m) {
				if dp[x][y] > grid[x][y]+curr.cost {
					dp[x][y] = grid[x][y] + curr.cost
					heap.Push(&mh, P{x, y, curr.cost + grid[x][y]})
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minimum_obstacle_removal_to_reach_corner([][]int{
		{0, 1, 1}, {1, 1, 0}, {1, 1, 0},
	}))
}
