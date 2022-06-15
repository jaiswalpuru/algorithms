package main

import "fmt"
import "math"

//--------------------brute force----------------------
func longest_string_chain(words []string) int {
	res := math.MinInt64
	n := len(words)

	exists := make(map[string]bool)
	for i := 0; i < n; i++ {
		exists[words[i]] = true
	}
	for i := 0; i < n; i++ {
		res = max(res, dfs(words, words[i], exists))
	}
	return res
}

func dfs(words []string, word string, exists map[string]bool) int {
	res := 1
	for i := 0; i < len(word); i++ {
		new_word := word[0:i] + word[i+1:]
		if exists[new_word] {
			cur_len := 1 + dfs(words, new_word, exists)
			res = max(res, cur_len)
		}
	}
	return res
}

//--------------------brute force----------------------

//-----------------------memoization----------------------
func longest_string_chain_memo(words []string) int {
	n := len(words)

	exists := make(map[string]bool)
	for i := 0; i < n; i++ {
		exists[words[i]] = true
	}

	res := 0
	memo := make(map[string]int)
	for i := 0; i < n; i++ {
		res = max(res, dfs_eff(words, words[i], exists, memo))
	}
	return res
}

func dfs_eff(words []string, word string, exists map[string]bool, memo map[string]int) int {
	if _, ok := memo[word]; ok {
		return memo[word]
	}

	res := 1
	for i := 0; i < len(word); i++ {
		new_word := word[0:i] + word[i+1:]
		if exists[new_word] {
			res = max(res, 1+dfs_eff(words, new_word, exists, memo))
		}
	}
	memo[word] = res
	return res
}

//-----------------------memoization----------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longest_string_chain_memo([]string{
		"a", "b", "ba", "bca", "bda", "bdca",
	}))
}
