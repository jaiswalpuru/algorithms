package main

import (
	"fmt"
	"strconv"
)

func ctos(arr []int) string { return strconv.Itoa(arr[0]) + "->" + strconv.Itoa(arr[1]) }

func make_graph(coords [][]int) (map[int][][]int, map[int][][]int) {
	x_map := make(map[int][][]int)
	y_map := make(map[int][][]int)

	n := len(coords)
	for i := 0; i < n; i++ {
		x_map[coords[i][0]] = append(x_map[coords[i][0]], coords[i])
		y_map[coords[i][1]] = append(y_map[coords[i][1]], coords[i])
	}

	return x_map, y_map
}

func stones(stones [][]int) int {
	xmp, ymp := make_graph(stones)
	cnt := 0
	visited := make(map[string]bool)
	n := len(stones)
	for i := 0; i < n; i++ {
		if !visited[ctos(stones[i])] {
			q := [][]int{stones[i]}
			for len(q) > 0 {
				curr := q[0]
				q = q[1:]
				cnt += 1
				x, y := curr[0], curr[1]
				visited[ctos(curr)] = true
				for j := 0; j < len(xmp[x]); j++ {
					if !visited[ctos(xmp[x][j])] {
						q = append(q, xmp[x][j])
						visited[ctos(xmp[x][j])] = true
					}
				}
				for j := 0; j < len(ymp[y]); j++ {
					if !visited[ctos(ymp[y][j])] {
						q = append(q, ymp[y][j])
						visited[ctos(ymp[y][j])] = true
					}
				}
			}
			cnt -= 1
		}
	}
	return cnt
}

//--------------------------------dfs------------------------------------------
func remove_stone_dfs(stones [][]int) int {
	x_map, y_map := make_graph(stones)
	cnt := 0
	visited := make(map[string]bool)
	n := len(stones)
	for i := 0; i < n; i++ {
		if !visited[ctos(stones[i])] {
			cnt++
			dfs(stones, &visited, stones[i], x_map, y_map)
		}
	}
	return len(stones) - cnt
}

func dfs(stones [][]int, visited *map[string]bool, stone []int, x_map, y_map map[int][][]int) {
	(*visited)[ctos(stone)] = true
	for i := 0; i < len(x_map[stone[0]]); i++ {
		if !(*visited)[ctos(x_map[stone[0]][i])] {
			dfs(stones, visited, x_map[stone[0]][i], x_map, y_map)
		}
	}
	for i := 0; i < len(y_map[stone[1]]); i++ {
		if !(*visited)[ctos(x_map[stone[1]][i])] {
			dfs(stones, visited, x_map[stone[1]][i], x_map, y_map)
		}
	}
}

//--------------------------------dfs------------------------------------------

func main() {
	fmt.Println(remove_stone_dfs([][]int{
		{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2},
	}))
}
