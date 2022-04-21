package main

import "fmt"

type TrieNode struct {
	is_word  bool
	children map[byte]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func New() Trie { return Trie{root: &TrieNode{children: make(map[byte]*TrieNode)}} }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (this *Trie) insert(word string) {
	root := this.root

	for i := 0; i < len(word); i++ {
		index := word[i]
		if root.children[index] == nil {
			root.children[index] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		root = root.children[index]
	}
	root.is_word = true
}

func (this *Trie) longest_prefix(word string) int {
	l := 0
	root := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if root.children == nil || root.children[ch] == nil {
			break
		}

		root = root.children[ch]
		if root.is_word {
			l = i + 1
		}
	}
	return l
}

func bold_words_in_string(words []string, s string) string {
	n := len(words)
	root := New()
	for i := 0; i < n; i++ {
		root.insert(words[i])
	}

	res := []byte{}
	rem := 0
	is_open := false
	for i := 0; i < len(s); i++ {
		pref := s[i:]
		rem = max(root.longest_prefix(pref), rem)

		if !is_open && rem > 0 {
			is_open = true
			res = append(res, []byte{'<', 'b', '>'}...)
		} else if is_open && rem <= 0 {
			is_open = false
			res = append(res, []byte{'<', '/', 'b', '>'}...)
		}
		res = append(res, s[i])
		rem--
	}

	if is_open {
		res = append(res, []byte{'<', '/', 'b', '>'}...)
	}
	return string(res)
}

func main() {
	fmt.Println(bold_words_in_string([]string{"ab", "bc"}, "aabcd"))
}
