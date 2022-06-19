package main

import "fmt"
import "sort"

type TrieNode struct {
	words    []string
	children []*TrieNode
}

func New() *TrieNode { return &TrieNode{words: make([]string, 0), children: make([]*TrieNode, 26)} }

func insert(root *TrieNode, word string) {
	cur := root
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if cur.children[ch] == nil {
			cur.children[ch] = New()
		}
		if len(cur.children[ch].words) < 3 {
			cur.children[ch].words = append(cur.children[ch].words, word)
		}
		cur = cur.children[ch]
	}
}

func search_word(str string, root *TrieNode) [][]string {
	words := [][]string{}
	curr := root
	for i := 0; i < len(str); i++ {
		ch := str[i] - 'a'
		if curr.children[ch] != nil {
			words = append(words, curr.children[ch].words)
			curr = curr.children[ch]
		} else {
			words = append(words, []string{})
			curr.children[ch] = New()
			curr = curr.children[ch]
		}
	}

	return words
}

func suggested_products(prod []string, search string) [][]string {
	sort.Strings(prod)
	tr := New()
	for i := 0; i < len(prod); i++ {
		insert(tr, prod[i])
	}

	return search_word(search, tr)
}

func main() {
	fmt.Println(suggested_products([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
}
