/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells,
where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

*/

package main

import "fmt"

var (
	row, col int
	visited  = [][]bool{}
)

func dfs(mat [][]byte, word string, index, i, j int) bool {

	//if the index has the reached the length of the word
	if index == len(word) {
		return true
	}

	//edge cases check
	if i < 0 || j < 0 || i >= row || j >= col {
		return false
	}

	//already visited path
	if visited[i][j] {
		return false
	}

	//doesn't match
	if mat[i][j] != word[index] {
		return false
	}

	visited[i][j] = true
	res := dfs(mat, word, index+1, i-1, j) ||
		dfs(mat, word, index+1, i+1, j) ||
		dfs(mat, word, index+1, i, j+1) ||
		dfs(mat, word, index+1, i, j-1)

	visited[i][j] = false

	return res
}

func check_word(mat [][]byte, word string) bool {

	row, col = len(mat), len(mat[0])

	visited = make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(mat, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func main() {
	matrix := [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}

	fmt.Println(check_word(matrix, "abcced"))
	fmt.Println(check_word(matrix, "see"))
	fmt.Println(check_word(matrix, "abcb"))
}
