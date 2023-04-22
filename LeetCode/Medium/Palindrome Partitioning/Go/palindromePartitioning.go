package main

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	res := [][]string{}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	backtrack(s, &visited, 0, &res, []string{})
	return res
}

func backtrack(s string, visited *[][]bool, start int, res *[][]string, temp []string) {
	if start == len(s) {
		*res = append(*res, append([]string{}, temp...))
		return
	}
	for end := start; end < len(s); end++ {
		if s[start] == s[end] && (end-start < 2 || (*visited)[start+1][end-1]) {
			(*visited)[start][end] = true
			temp = append(temp, s[start:end+1])
			backtrack(s, visited, end+1, res, temp)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(partition("aab"))
}
