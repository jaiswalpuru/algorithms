package main

import "fmt"

type Trie struct {
	children [26]*Trie
	word     string
}

func New() *Trie { return &Trie{} }

func (this *Trie) Insert(word string) {
	root := this
	for _, ch := range word {
		index := ch - 'a'
		if root.children[index] == nil {
			root.children[index] = &Trie{}
		}
		root = root.children[index]
	}
	root.word = word
}

//----------------------------backtrack brute force way (TLE)------------------------------------------
func word_search_brute(grid [][]byte, words []string) []string {
	res := []string{}
	n, m := len(grid), len(grid[0])

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for _, word := range words {
		f := false
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] == word[0] {
					if _dfs(grid, word, i, j, 0, &visited) {
						res = append(res, word)
						f = true
						break
					}
				}
			}
			if f {
				break
			}
		}
	}

	return res
}

func _dfs(grid [][]byte, word string, i, j, k int, visited *[][]bool) bool {

	if k == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || (*visited)[i][j] {
		return false
	}

	if grid[i][j] != word[k] {
		return false
	}

	(*visited)[i][j] = true
	res := _dfs(grid, word, i+1, j, k+1, visited) || _dfs(grid, word, i-1, j, k+1, visited) || _dfs(grid, word, i, j+1, k+1, visited) || _dfs(grid, word, i, j-1, k+1, visited)
	(*visited)[i][j] = false
	return res

}

//------------------------------------------------------------------------------------

//---------------------------------------optimized---------------------------------------------
func word_search(grid [][]byte, words []string) []string {
	res := []string{}
	n, m := len(grid), len(grid[0])

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	trie := New()
	for _, word := range words {
		trie.Insert(word)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(grid, trie, i, j, &res, &visited)
		}
	}

	return res
}

func dfs(grid [][]byte, root *Trie, i, j int, res *[]string, visited *[][]bool) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || (*visited)[i][j] {
		return
	}

	temp := grid[i][j]
	if root.children[temp-'a'] == nil {
		return
	}

	root = root.children[temp-'a']
	if len(root.word) > 0 {
		*res = append(*res, root.word)
		root.word = ""
	}

	(*visited)[i][j] = true
	dfs(grid, root, i+1, j, res, visited)
	dfs(grid, root, i-1, j, res, visited)
	dfs(grid, root, i, j+1, res, visited)
	dfs(grid, root, i, j-1, res, visited)
	(*visited)[i][j] = false
}

//---------------------------------------------------------------------------------------------

func main() {
	fmt.Println(word_search([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain"}))
}
