/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are
horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example 1:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Example 2:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Example 3:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
*/

package main

import "fmt"

var (
	row, col int
	visited  = [][]bool{}
)

func dfs(matrix [][]byte, word string, idx, i, j int) bool {
	if idx == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= row || j >= col {
		return false
	}

	if visited[i][j] {
		return false
	}

	if matrix[i][j] != word[idx] {
		return false
	}

	visited[i][j] = true
	res := dfs(matrix, word, idx+1, i+1, j) ||
		dfs(matrix, word, idx+1, i-1, j) ||
		dfs(matrix, word, idx+1, i, j+1) ||
		dfs(matrix, word, idx+1, i, j-1)
	visited[i][j] = false

	return res

}

func word_search(matrix [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}

	row, col = len(matrix), len(matrix[0])

	visited = make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(matrix, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(word_search([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))

}
