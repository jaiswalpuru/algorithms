/*
	Implement an autocomplete system. That is, given a query string s and a set of all possible query strings,
	return all strings in the set that have s as a prefix.
	For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
	Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.
*/

package main

type trie struct {
	letter   rune
	children []*trie
	is_leaf  bool
	metadata string
}

func new_trie() *trie { return &trie{children: make([]*trie, 0), letter: '#', metadata: ""} }

func (t *trie) has_child(letter rune) (bool, *trie) {

	for _, child := range t.children {
		if child.letter == letter {
			return true, child
		}
	}
	return false, nil

}

func (t *trie) add_child(letter rune) *trie {
	nt := new_trie()
	nt.letter = letter
	t.children = append(t.children, nt)
	return nt
}

func (t *trie) add_word(word, m string) *trie {

	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {

		if exists, val := node.has_child(letters[i]); exists {
			node = val
		} else {
			node = node.add_child(letters[i])
		}

		i++
		if i == n {
			node.is_leaf = true
			node.metadata = m
		}
	}

	return node
}

func (t *trie) find_node(word string) *trie {

	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {
		if exists, val := node.has_child(letters[i]); exists {
			node = val
		} else {
			return nil
		}
		i++
	}

	return node
}

func (t *trie) find(word string) (*trie, string) {

	node := t.find_node(word)

	if node == nil {
		return nil, ""
	} else {
		return node, node.metadata
	}

}

var words = make([]string, 0)

func elements(t *trie) {
	for _, val := range t.children {
		if val.is_leaf {
			words = append(words, val.metadata)
		}
		elements(val)
	}
}

func main() {

	queryString := "d"
	arr := []string{"dog", "deer", "deal"}

	nt := new_trie()
	for i := 0; i < len(arr); i++ {
		nt.add_word(arr[i], arr[i])
	}

	node, _ := nt.find(queryString)

	if node != nil {
		if node.is_leaf {
			words = append(words, node.metadata)
		}

		for _, val := range node.children {
			if val.is_leaf {
				words = append(words, val.metadata)
			}
			elements(val)
		}

	}
}
