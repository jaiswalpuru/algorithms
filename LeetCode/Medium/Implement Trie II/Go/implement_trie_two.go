package main

type TrieNode struct {
	is_word  bool
	cnt      int
	children [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie { return Trie{root: &TrieNode{}} }

func (this *Trie) Insert(word string) {
	root := this.root
	for _, ch := range word {
		index := ch - 'a'
		if root.children[index] == nil {
			root.children[index] = &TrieNode{}
		}
		root = root.children[index]
	}
	root.is_word = true
	root.cnt++
}

func (this *Trie) CountWordsEqualTo(word string) int {
	root := this.root
	for _, ch := range word {
		index := ch - 'a'
		if root.children[index] == nil {
			return 0
		}
		root = root.children[index]
	}

	if root.is_word {
		return root.cnt
	}

	return 0
}

func (this *Trie) Erase(word string) {
	root := this.root
	for _, ch := range word {
		index := ch - 'a'
		if root.children[index] == nil {
			break
		}
		root = root.children[index]
	}
	root.cnt--
	if root.cnt == 0 {
		root.is_word = false
	}
}

func (this *Trie) CountWordsStartingWith(prefix string) int {
	root := this.root
	count := 0
	for _, ch := range prefix {
		index := ch - 'a'
		if root.children[index] == nil {
			return 0
		}
		root = root.children[index]
	}
	count += root.cnt
	cnt_word(root, &count)
	return count
}

func cnt_word(root *TrieNode, count *int) {
	for i := 0; i < len(root.children); i++ {
		if root.children[i] != nil {
			if root.children[i].is_word {
				*count += root.children[i].cnt
			}
			cnt_word(root.children[i], count)
		}
	}
}

func main() {}
