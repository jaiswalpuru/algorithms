/*
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.

A palindrome string is a string that reads the same backward as forward.
*/

package main

import "fmt"

func is_palindrome(str string, l, h int) bool {
	for l < h {
		if str[l] != str[h] {
			return false
		}
		l++
		h--
	}
	return true
}

func dfs(str string, res *[][]string, cur_list []string, idx int) {
	if idx >= len(str) {
		*res = append(*res, append([]string{}, cur_list...))
	}

	for end := idx; end < len(str); end++ {
		if is_palindrome(str, idx, end) {
			temp := append(cur_list, str[idx:end+1])
			dfs(str, res, temp, end+1)
			temp = temp[:len(temp)-1]
		}
	}
}

func partition(str string) [][]string {
	res := make([][]string, 0)
	current_list := make([]string, 0)
	dfs(str, &res, current_list, 0)
	return res
}

//partitioning backtracking with DP
func partition_dp_bt(str string) [][]string {
	n := len(str)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}

	res := make([][]string, 0)
	cur_list := make([]string, 0)
	dfs_dp(str, &res, cur_list, &visited, 0)
	return res
}

func dfs_dp(str string, res *[][]string, cur_list []string, visited *[][]bool, start int) {
	if start >= len(str) {
		*res = append(*res, append([]string{}, cur_list...))
	}

	for end := start; end < len(str); end++ {
		if (str[start] == str[end]) && (end-start <= 2 || (*visited)[start+1][end-1]) {
			(*visited)[start][end] = true
			temp := append(cur_list, str[start:end+1])
			dfs_dp(str, res, temp, visited, end+1)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(partition_dp_bt("aab"))
}
