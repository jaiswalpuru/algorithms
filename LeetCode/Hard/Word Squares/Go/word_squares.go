package main

import "fmt"

type Trie struct {
	children [26]*Trie
	word_ind []int
}

func New() *Trie { return &Trie{} }

func (this *Trie) Insert(word string, ind int) {
	root := this
	for i := 0; i < len(word); i++ {
		ch := int(word[i]) - 'a'
		if root.children[ch] == nil {
			root.children[ch] = New()
		}
		root = root.children[ch]
		root.word_ind = append(root.word_ind, ind)
	}
}

func (this *Trie) Search(word string) []int {
	root := this
	for i := 0; i < len(word); i++ {
		ch := int(word[i]) - int('a')
		if root.children[ch] == nil {
			return nil
		}
		root = root.children[ch]
	}
	return root.word_ind
}

func word_squares(words []string) [][]string {
	res := [][]string{}
	trie := New()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i], i)
	}
	n := len(words[0])
	for i := 0; i < len(words); i++ {
		temp := []string{words[i]}
		backtrack(temp, 1, words, &res, n, trie)
	}
	return res
}

func backtrack(set []string, ind int, words []string, res *[][]string, n int, trie *Trie) {
	if ind == n {
		*res = append(*res, append([]string{}, set...))
		return
	}
	prefix := ""
	for i := 0; i < len(set); i++ {
		prefix += string(set[i][ind])
	}
	pref_words := get_all_pref_words(prefix, trie)
	for i := 0; i < len(pref_words); i++ {
		set = append(set, words[pref_words[i]])
		backtrack(set, ind+1, words, res, n, trie)
		set = set[:len(set)-1]
	}
}

func get_all_pref_words(prefix string, trie *Trie) []int {
	root := trie
	for i := 0; i < len(prefix); i++ {
		ch := int(prefix[i]) - int('a')
		if root.children[ch] == nil {
			return nil
		}
		root = root.children[ch]
	}
	return root.word_ind
}

func main() {
	fmt.Println(word_squares([]string{"area", "lead", "wall", "lady", "ball"}))
}
