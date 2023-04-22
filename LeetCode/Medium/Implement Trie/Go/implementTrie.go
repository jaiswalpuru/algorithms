package main

type TrieNode struct {
	isWord   bool
	children [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	root := this.root
	size := len(word)
	for i := 0; i < size; i++ {
		ch := int(word[i] - 'a')
		if root.children[ch] == nil {
			root.children[ch] = &TrieNode{}
		}
		root = root.children[ch]
	}
	root.isWord = true
}

func (this *Trie) Search(word string) bool {
	root := this.root
	size := len(word)
	for i := 0; i < size; i++ {
		ch := int(word[i] - 'a')
		if root.children[ch] == nil {
			return false
		}
		root = root.children[ch]
	}
	if root.isWord {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	root := this.root
	size := len(prefix)
	for i := 0; i < size; i++ {
		ch := int(prefix[i] - 'a')
		if root.children[ch] == nil {
			return false
		}
		root = root.children[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
