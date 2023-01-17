package main

import "fmt"

type Trie struct {
	children [26]*Trie
	word     string
}

func New() *Trie { return &Trie{} }
func (this *Trie) Insert(word string, prefix string) string {
	root := this
	f := true
	k := 0
	for i := 0; i < len(word); i++ {
		ch := int(word[i] - 'a')
		if f {
			if k < len(prefix) && word[i] == prefix[k] {
				k++
			} else {
				f = false
			}
		}
		if root.children[ch] == nil {
			root.children[ch] = New()
		}
		root = root.children[ch]
	}
	root.word = word
	return prefix[:k]
}

func longest_common_prefix(strs []string) string {
	trie := New()
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		word := trie.Insert(strs[i], prefix)
		if len(word) < len(prefix) {
			prefix = word
		}
	}
	return prefix
}

func main() {
	fmt.Println(longest_common_prefix([]string{"flower", "flow", "flight"}))
}
