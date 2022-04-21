package main

import "fmt"

type TrieNode struct {
	is_word  bool
	children map[byte]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie { return Trie{root: &TrieNode{children: make(map[byte]*TrieNode)}} }

func (this *Trie) insert(word string) {
	root := this.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if root.children[ch] == nil {
			root.children[ch] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		root = root.children[ch]
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

func add_bold_tags(s string, words []string) string {
	trie := Constructor()
	n := len(words)
	for i := 0; i < n; i++ {
		trie.insert(words[i])
	}

	rem := 0
	is_open := false
	res := []byte{}

	for i := 0; i < len(s); i++ {
		rem = max(trie.longest_prefix(s[i:]), rem)
		if !is_open && rem > 0 {
			is_open = true
			res = append(res, []byte{'<', 'b', '>'}...)
		} else if rem <= 0 && is_open {
			is_open = false
			res = append(res, []byte{'<', '/', 'b', '>'}...)
		}
		rem--
		res = append(res, s[i])
	}

	if is_open {
		res = append(res, []byte{'<', '/', 'b', '>'}...)
	}
	return string(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(add_bold_tags("abcxyz123", []string{"abc", "123"}))
}
