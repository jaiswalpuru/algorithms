package main

import "fmt"

func concatenated_words(words []string) []string {
	mp := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		mp[words[i]] = true
	}
	res := []string{}
	visited := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		dfs(mp, &res, 0, words[i], 0, &visited)
	}
	return res
}

func dfs(mp map[string]bool, res *[]string, cnt int, word string, ind int, visited *map[string]bool) {
	if ind == len(word) {
		if cnt > 1 {
			if !(*visited)[word] {
				*res = append(*res, word)
				(*visited)[word] = true
			}
		}
		return
	}
	s := ""
	for i := ind; i < len(word); i++ {
		s += string(word[i])
		if mp[s] {
			dfs(mp, res, cnt+1, word, i+1, visited)
		}
	}
}

func main() {
	fmt.Println(concatenated_words([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}
