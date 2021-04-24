// Implement an autocomplete system. That is, given a query string s and a set of all possible query strings,
// return all strings in the set that have s as a prefix.
// For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
// Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.

package main

import "fmt"

type trie struct {
	letter   rune
	children []*trie
	isLeaf   bool
	metaData string
}

func newTrie() *trie {

	nt := &trie{}
	nt.children = make([]*trie, 0)
	nt.letter = '#'
	nt.metaData = ""
	return nt
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

func (t *trie) addWord(word, m string) *trie {

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
			node.metaData = m
		}
	}

	return node
}

func (t *trie) findNode(word string) *trie {

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
	node := t.findNode(word)
	if node == nil {
		return nil, ""
	} else {
		return node, node.metaData
	}
}

var words = make([]string, 0)

func elements(t *trie) {
	for _, val := range t.children {
		if val.isLeaf {
			words = append(words, val.metaData)
		} else {
			elements(val)
		}
	}
}

func main() {

	queryString := "d"
	arr := []string{"dog", "deer", "deal"}

	nt := newTrie()
	for i := range arr {
		nt.addWord(arr[i], arr[i])
	}

	//here node contains the last letter which is pointed by the prefix
	node, _ := nt.find(queryString)

	//what need to be done here is after receiving node we need to process all the words which are accessible
	//from this node
	if node != nil {

		if node.isLeaf {
			words = append(words, node.metaData)
		}

		for _, val := range node.children {
			if val.isLeaf {
				words = append(words, val.metaData)
			}
			elements(val)
		}

	}
	fmt.Println(words)
}
