package main

type TrieNode struct {
	children [26]*TrieNode
	is_word  bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	root := this.root
	for _, ch := range word {
		index := ch - 'a'
		if root.children[index] == nil {
			root.children[index] = &TrieNode{}
		}
		root = root.children[index]
	}
	root.is_word = true
}

func (this *WordDictionary) Search(word string) bool {
	return search(word, this.root)
}

func search(word string, root *TrieNode) bool {
	for k, ch := range word {
		if ch == '.' {
			for i := 0; i < len(root.children); i++ {
				if root.children[i] != nil {
					if _search(word[k+1:], root.children[i]) {
						return true
					}
				}
			}
			return false
		} else {
			index := ch - 'a'
			if root.children[index] == nil {
				return false
			}
			root = root.children[index]
		}
	}
	return root.is_word
}

func main() {}
