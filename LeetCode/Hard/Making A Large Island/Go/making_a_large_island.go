package main

import "fmt"

//----------------------brute force----------------------
func making_a_large_island(arr [][]int) int {
	res := 0
	n, m := len(arr), len(arr[0])
	zero := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				zero = true
				arr[i][j] = 1
				t := 0
				visited := make([][]bool, n)
				for i := 0; i < n; i++ {
					visited[i] = make([]bool, m)
				}
				dfs(i, j, arr, &t, visited)
				res = max(res, t)
				arr[i][j] = 0
			}
		}
	}
	if zero {
		return res
	}
	return n * n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dir = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

func dfs(i, j int, arr [][]int, t *int, visited [][]bool) {
	visited[i][j] = true
	*t++
	for k := 0; k < 4; k++ {
		x, y := dir[k][0]+i, dir[k][1]+j
		if !(x < 0 || y < 0 || x >= len(arr) || y >= len(arr[0]) || arr[x][y] == 0 || visited[x][y]) {
			dfs(x, y, arr, t, visited)
		}
	}
}

//----------------------brute force----------------------

//----------------------brute force 2----------------------
func making_a_large_island_bf(arr [][]int) int {
	n := len(arr)
	zero := false
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				zero = true
				arr[i][j] = 1
				res = max(res, _dfs(arr, i, j))
				arr[i][j] = 0
			}
		}
	}
	if zero {
		return res
	}
	return n * n
}

func _dfs(arr [][]int, r, c int) int {
	seen := make(map[int]bool)
	st := []int{r*len(arr) + c}
	seen[r*len(arr)+c] = true
	for len(st) > 0 {
		curr := st[len(st)-1]
		st = st[:len(st)-1]
		i, j := curr/len(arr), curr%len(arr)
		for k := 0; k < 4; k++ {
			nr, nc := i+dir[k][0], j+dir[k][1]
			if !seen[nr*len(arr)+nc] && nr >= 0 && nr < len(arr) && nc >= 0 && nc < len(arr) && arr[nr][nc] == 1 {
				st = append(st, nr*len(arr)+nc)
				seen[nr*len(arr)+nc] = true
			}
		}
	}
	return len(seen)
}

//----------------------brute force 2----------------------

//----------------------efficient way----------------------
func making_a_large_island_eff(arr [][]int) int {
	n := len(arr)
	ind := 2
	area := make([]int, n*n+2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 1 {
				area[ind] = dfs_eff(arr, i, j, ind)
				ind++
			}
		}
	}
	ans := 0
	for i := 0; i < len(area); i++ {
		ans = max(ans, area[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				seen := make(map[int]bool)
				st := neighbors(i, j, n)
				for x := range st {
					if arr[st[x]/n][st[x]%n] > 1 {
						seen[arr[st[x]/n][st[x]%n]] = true
					}
				}
				bns := 1
				for k := range seen {
					bns += area[k]
				}
				ans = max(ans, bns)
			}
		}
	}
	return ans
}

func dfs_eff(arr [][]int, r, c int, ind int) int {
	ans := 1
	arr[r][c] = ind
	st := neighbors(r, c, len(arr))
	for x := range st {
		if arr[st[x]/len(arr)][st[x]%len(arr)] == 1 {
			arr[st[x]/len(arr)][st[x]%len(arr)] = ind
			ans += dfs_eff(arr, st[x]/len(arr), st[x]%len(arr), ind)
		}
	}
	return ans
}

func neighbors(i, j int, n int) []int {
	neigh := []int{}
	for k := 0; k < 4; k++ {
		r, c := i+dir[k][0], j+dir[k][1]
		if r >= 0 && c >= 0 && r < n && c < n {
			neigh = append(neigh, r*n+c)
		}
	}
	return neigh
}

//----------------------efficient way----------------------

func main() {
	fmt.Println(making_a_large_island_eff([][]int{
		{1, 1}, {1, 0},
	}))
}
