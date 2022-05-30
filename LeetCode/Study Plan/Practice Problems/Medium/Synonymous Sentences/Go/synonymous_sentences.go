package main

import "fmt"
import "strings"
import "sort"

func synonymous_sentences(syn [][]string, text string) []string {
	mp := make(map[string][]string)
	res := []string{}
	for i := 0; i < len(syn); i++ {
		mp[syn[i][0]] = append(mp[syn[i][0]], syn[i][1])
		mp[syn[i][1]] = append(mp[syn[i][1]], syn[i][0])
	}

	text_str := strings.Split(text, " ")
	_backtrack(text_str, &res, mp, 0, []string{})
	return res
}

func _backtrack(text_str []string, res *[]string, mp map[string][]string, ind int, s []string) {
	if ind >= len(text_str) {
		*res = append(*res, strings.Join(s, " "))
		return
	}

	word := text_str[ind]
	if mp[word] == nil {
		temp := append(s, word)
		_backtrack(text_str, res, mp, ind+1, temp)
		temp = temp[:len(temp)-1]
	} else {
		synon := []string{}
		visited := make(map[string]bool)
		dfs(word, &synon, mp, &visited)
		sort.Strings(synon)
		for i := 0; i < len(synon); i++ {
			temp := append(s, synon[i])
			_backtrack(text_str, res, mp, ind+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
}

func dfs(word string, synon *[]string, mp map[string][]string, visited *map[string]bool) {
	if (*visited)[word] {
		return
	}
	(*visited)[word] = true
	for i := 0; i < len(mp[word]); i++ {
		dfs(mp[word][i], synon, mp, visited)
	}
	(*synon) = append(*synon, word)
}

func main() {
	fmt.Println(synonymous_sentences([][]string{
		{"happy", "joy"}, {"sad", "sorrow"}, {"joy", "cheerful"},
	}, "I am happy today but was sad yesterday"))
}
