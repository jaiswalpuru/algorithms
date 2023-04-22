package main

import "fmt"

//------------------using union find---------------------
func num_provinces_union_find(is_connected [][]int) int {
	n := len(is_connected)
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if is_connected[i][j] == 1 {
				union(i, j, &parent, &size)
			}
		}
	}
	provinces := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			provinces++
		}
	}
	return provinces
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] > (*size)[y] {
		(*parent)[y] = x
		(*size)[x] += (*size)[y]
	} else {
		(*parent)[x] = y
		(*size)[y] += (*size)[x]
	}
}

func find(x int, parent *[]int) int {
	if (*parent)[x] != x {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
}

//------------------using union find---------------------

//------------------using dfs---------------------
func num_provinces(is_connected [][]int) int {
	n := len(is_connected)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if is_connected[i][j] == 1 && visited[i][j] != true {
				cnt++
				visited[i][j] = true
				dfs(i, j, is_connected, &visited)
			}
		}
	}
	return cnt
}

func dfs(i, j int, is_connected [][]int, visited *[][]bool) {
	(*visited)[i][j] = true
	for k := 0; k < len(is_connected); k++ {
		if !(*visited)[j][k] && is_connected[j][k] == 1 {
			dfs(j, k, is_connected, visited)
		}
	}
}

//------------------dfs------------------------------

func main() {
	fmt.Println(num_provinces([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))
}
