/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or
 vertically neighboring.
The same letter cell may not be used more than once.
*/

package main

import "fmt"

var visited [][]bool

func _word_search(grid [][]byte, word string, i, j, k, n, m int) bool {
	if k == len(word) {
		return true
	}

	if i >= n || j >= m || i < 0 || j < 0 {
		return false
	}

	if visited[i][j] {
		return false
	}

	if grid[i][j] != word[k] {
		return false
	}

	visited[i][j] = true
	res := _word_search(grid, word, i+1, j, k+1, n, m) ||
		_word_search(grid, word, i-1, j, k+1, n, m) ||
		_word_search(grid, word, i, j+1, k+1, n, m) ||
		_word_search(grid, word, i, j-1, k+1, n, m)

	visited[i][j] = false
	return res
}

func word_search(grid [][]byte, word string) bool {
	n, m := len(grid), len(grid[0])
	visited = make([][]bool, n)

	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _word_search(grid, word, i, j, 0, n, m) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(word_search([][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}, "ABCCED"))
}
