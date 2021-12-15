package main

import (
	"fmt"
	"strconv"
)

type trie struct {
	letter   rune
	children []*trie
	isleaf   bool
	meta     map[string]string
}

func (t *trie) hasChild(a rune) (bool, *trie) {
	for _, child := range t.children {
		if child.letter == a {
			return true, child
		}
	}
	return false, nil
}

func newTrie() *trie {
	nt := &trie{}
	nt.children = []*trie{}
	nt.letter = '#'
	nt.meta = make(map[string]string)
	return nt
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
			node.isleaf = true
			node.meta[word] = m
		}
	}
	return node
}

func (t *trie) findNode(word string) *trie {
	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {
		if exist, value := node.hasChild(letters[i]); exist {
			node = value
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
		return node, node.meta[word]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	nw := newTrie()
	for i := 0; i < n; i++ {
		var b string
		var a int
		fmt.Scan(&b, &a)
		nw.addWord(b, string(rune(a)))
	}

	var s string
	fmt.Scan(&s)
	ans := ""
	i, j := 0, 1
	for i = 0; i < len(s); {
		if j > len(s) {
			if i == 0 {
				ans = "DECODE FAIL AT INDEX " + strconv.Itoa(0)
			} else {
				ans = "DECODE FAIL AT INDEX " + strconv.Itoa(j-2)
			}
			break
		}
		t, val := nw.find(s[i:j])
		if t != nil && t.isleaf {
			ans += val
			i = j
			j++
		} else {
			j++
		}
	}
	fmt.Println(ans)
}
