package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	children [26]*Trie
	word     string
}

func New() *Trie { return &Trie{} }

func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		index := int(word[i]) - int('a')
		if root.children[index] == nil {
			root.children[index] = New()
		}
		root = root.children[index]
	}
	root.word = word
}

func (this *Trie) Search(word string) string {
	root := this
	for i := 0; i < len(word); i++ {
		index := int(word[i]) - int('a')
		if len(root.word) > 0 {
			return root.word
		}
		if root.children[index] == nil {
			return ""
		}
		root = root.children[index]
	}
	return root.word
}

func replace_words(dict []string, sen string) string {
	trie := New()
	for i := 0; i < len(dict); i++ {
		trie.Insert(dict[i])
	}
	str := strings.Split(sen, " ")
	for i := 0; i < len(str); i++ {
		curr := str[i]
		v := trie.Search(curr)
		if len(v) > 0 {
			str[i] = v
		}
	}
	return strings.Join(str, " ")
}

func main() {
	fmt.Println(replace_words([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
