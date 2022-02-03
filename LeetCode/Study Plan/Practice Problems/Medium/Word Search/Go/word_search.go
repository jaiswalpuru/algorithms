package main

import "fmt"

func search(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	word_len := len(word)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(board, word, i, j, 0, word_len, &visited) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k, word_len int, visited *[][]bool) bool {

	if k == word_len {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if (*visited)[i][j] {
		return false
	}
	if word[k] != board[i][j] {
		return false
	}

	(*visited)[i][j] = true
	res := dfs(board, word, i+1, j, k+1, word_len, visited) || dfs(board, word, i, j+1, k+1, word_len, visited) ||
		dfs(board, word, i, j-1, k+1, word_len, visited) || dfs(board, word, i-1, j, k+1, word_len, visited)
	(*visited)[i][j] = false
	return res
}

func main() {
	fmt.Println(search([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE"))
}
