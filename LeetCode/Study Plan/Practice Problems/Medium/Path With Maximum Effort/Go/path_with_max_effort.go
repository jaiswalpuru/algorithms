package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Triplet struct{ x, y, val int }

type MinHeap []Triplet

func (mh MinHeap) Len() int              { return len(mh) }
func (mh MinHeap) Less(i, j int) bool    { return mh[i].val < mh[j].val }
func (mh MinHeap) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(val interface{}) { *mh = append(*mh, val.(Triplet)) }
func (mh *MinHeap) Pop() interface{} {
	old := (*mh)[len(*mh)-1]
	(*mh) = (*mh)[:len(*mh)-1]
	return old
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

var (
	direction = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

//-----------------------------backtrack-----------------------------(TLE)
func heights(arr [][]int) int {
	n, m := len(arr)-1, len(arr[0])-1
	max_so_far := math.MaxInt64
	return _heights(arr, 0, 0, n, m, 0, &max_so_far)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func is_valid(x, y, n, m int) bool { return (x >= 0 && y >= 0 && x <= n && y <= m) }

func _heights(arr [][]int, i, j, n, m, max_diff int, max_so_far *int) int {
	if i == n && j == m {
		*max_so_far = min(*max_so_far, max_diff)
		return max_diff
	}

	curr_ht := arr[i][j]
	arr[i][j] = 0
	min_effort := math.MaxInt64
	for k := 0; k < 4; k++ {
		adj_x := i + direction[k][0]
		adj_y := j + direction[k][1]
		if is_valid(adj_x, adj_y, n, m) && arr[adj_x][adj_y] != 0 {
			curr_diff := abs(arr[adj_x][adj_y] - curr_ht)
			curr_max_diff := max(max_diff, curr_diff)
			if *max_so_far > curr_max_diff {
				res := _heights(arr, adj_x, adj_y, n, m, curr_max_diff, max_so_far)
				min_effort = min(min_effort, res)
			}
		}
	}
	arr[i][j] = curr_ht
	return min_effort
}

//--------------------------------------------------------------------------

func heights_dijkstra(arr [][]int) int {
	n, m := len(arr), len(arr[0])
	diff_matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		diff_matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			diff_matrix[i][j] = math.MaxInt64
		}
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	mh := &MinHeap{}
	diff_matrix[0][0] = 0
	heap.Push(mh, Triplet{x: 0, y: 0, val: diff_matrix[0][0]})

	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Triplet)
		x, y, val := curr.x, curr.y, curr.val
		visited[x][y] = true
		if x == n-1 && y == m-1 {
			return val
		}

		for i := 0; i < 4; i++ {
			adj_x := x + direction[i][0]
			adj_y := y + direction[i][1]
			if is_valid(adj_x, adj_y, n-1, m-1) && !visited[adj_x][adj_y] {
				curr_diff := abs(arr[x][y] - arr[adj_x][adj_y])
				max_diff := max(curr_diff, diff_matrix[x][y])
				if diff_matrix[adj_x][adj_y] > max_diff {
					diff_matrix[adj_x][adj_y] = max_diff
					heap.Push(mh, Triplet{x: adj_x, y: adj_y, val: max_diff})
				}
			}
		}
	}
	return diff_matrix[n-1][m-1]
}

func main() {
	fmt.Println(heights_dijkstra([][]int{
		{1, 2, 2},
		{3, 8, 2},
		{5, 3, 5},
	}))
}
