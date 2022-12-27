package main

import "fmt"

type Trie struct {
	children [26]*Trie
	word     string
}

func New() *Trie { return &Trie{} }

func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		ind := int(word[i]) - int('a')
		if root.children[ind] == nil {
			root.children[ind] = New()
		}
		root = root.children[ind]
	}
	root.word = word
}

func longest_word_with_all_prefixes(words []string) string {
	trie := New()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	res := ""
	for i := 0; i < 26; i++ {
		if trie.children[i] != nil {
			temp := ""
			dfs(trie.children[i], &temp)
			if len(temp) == len(res) {
				if temp < res {
					res = temp
				}
			} else {
				if len(temp) > len(res) {
					res = temp
				}
			}
		}
	}
	return res
}

func dfs(t *Trie, res *string) {
	if t.word == "" {
		return
	}
	if t.word != "" {
		if len(*res) == len(t.word) {
			if t.word < *res {
				*res = t.word
			}
		} else {
			if len(*res) < len(t.word) {
				*res = t.word
			}
		}
	}
	for i := 0; i < len(t.children); i++ {
		if t.children[i] != nil {
			dfs(t.children[i], res)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longest_word_with_all_prefixes([]string{"k", "ki", "kir", "kira", "kiran"}))
}
