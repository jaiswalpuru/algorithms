package main

import "fmt"

type trie struct {
	letter   rune
	children []*trie
	meta     map[string]string
	isLeaf   bool
}

func newTrie() *trie {
	return &trie{letter: '#', children: []*trie{}, meta: make(map[string]string)}
}

func (t *trie) hasChild(a rune) (bool, *trie) {
	for _, child := range t.children {
		if child.letter == a {
			return true, child
		}
	}
	return false, nil
}

func (t *trie) addChild(a rune) *trie {
	nt := newTrie()
	nt.letter = a
	t.children = append(t.children, nt)
	return nt
}

func (t *trie) addWord(word string) *trie {

	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {
		if exists, val := node.hasChild(letters[i]); exists {
			node = val
		} else {
			node = node.addChild(letters[i])
		}
		i++
		if i == n {
			node.isLeaf = true
			node.meta[word] = word
		}
	}
	return node
}

func (t *trie) findWord(word string) *trie {
	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {
		if exists, val := node.hasChild(letters[i]); exists {
			node = val
		} else {
			return nil
		}
		i++
	}
	return node
}

func (t *trie) find(word string) (*trie, string) {
	node := t.findWord(word)
	if node == nil {
		return nil, ""
	} else {
		return node, node.meta[word]
	}
}

func main() {
	nt := newTrie()
	nt.addWord("quick")
	nt.addWord("brown")
	nt.addWord("the")
	nt.addWord("fox")
	str := "thequickbrownfox"
	strRune := []rune(str)
	list := []string{}
	i := 0

	for j := 0; j < len(strRune); j++ {
		if _, val := nt.find(string(strRune[i : j+1])); val != "" {
			list = append(list, val)
			i = j + 1
		}
	}
	fmt.Println(list)
}
