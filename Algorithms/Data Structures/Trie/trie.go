package main

import "fmt"

/*
	Prefix tree or digital trees or mostly called trie
	It is a type of k-ary search tree.
	There are other forms of this data structure :
	Bitwise tries
	Compressed tries (used in web search engines for indexing.)
	Here I am using an array of 26 characters, assuming only lower case alphabets will be stored.
	If want to store unicode characters use map[byte]*Trie
*/

type Trie struct {
	children [26]*Trie
	word     string
}

func New() *Trie { return &Trie{} }

func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		ch := int(word[i] - 'a')
		if root.children[ch] == nil {
			root.children[ch] = New()
		}
		root = root.children[ch]
	}
	root.word = word
}

func (this *Trie) find(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		ch := int(word[i] - 'a')
		if root.children[ch] == nil {
			return false
		}
		root = root.children[ch]
	}
	return true
}

func main() {
	trie := New()
	trie.Insert("ball")
	trie.Insert("balloon")
	trie.Insert("zebra")
	fmt.Println("Search ball", trie.find("ball"))
	fmt.Println("Search zebra", trie.find("zebra"))
	fmt.Println("Search balloonm", trie.find("balloon"))
	fmt.Println("Search bat", trie.find("bat"))
}
