package main

import "fmt"

func findAllConcatenatedWordsInADict(words []string) []string {
	wordsMap := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]] = true
	}
	res := []string{}
	visited := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		dfs(words[i], 0, 0, &res, &visited, wordsMap)
	}
	return res
}

func dfs(word string, cnt, ind int, res *[]string, visited *map[string]bool, wordsMap map[string]bool) {
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
		if wordsMap[s] {
			dfs(word, cnt+1, i+1, res, visited, wordsMap)
		}
	}
}

func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}
