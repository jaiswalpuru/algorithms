/*
	Create PrefixMapSum class
	Implement a PrefixMapSum class with the following methods
	• insert(key string, val int): set a given key's value  in the map
	• sum (prefix string) return the sum of all values of keys
*/

package main

import "fmt"

type Trie struct {
	letter   rune
	children []*Trie
	is_leaf  bool
	metadata int
}

func new_trie() *Trie { return &Trie{letter: '#', children: make([]*Trie, 0), metadata: -1} }

func (t *Trie) has_child(letter rune) (bool, *Trie) {
	for _, child := range t.children {
		if child.letter == letter {
			return true, child
		}
	}
	return false, nil
}

func (t *Trie) find_node(word string) *Trie {

	letters, node, i := []rune(word), t, 0
	n := len(letters)

	for i < n {
		if exists, val := node.has_child(letters[i]); exists {
			node = val
		} else {
			return val
		}
		i++
	}

	return node
}

func (t *Trie) find(word string) (*Trie, int) {

	node := t.find_node(word)

	if node == nil {
		return nil, -1
	} else {
		return node, node.metadata
	}

}

func (t *Trie) add_child(letter rune) *Trie {
	nt := new_trie()
	nt.letter = letter
	t.children = append(t.children, nt)
	return nt
}

func (t *Trie) insert(word string, metadata int) {
	letter, node, i := []rune(word), t, 0
	n := len(letter)

	for i < n {
		if exists, val := node.has_child(letter[i]); exists {
			node = val
		} else {
			node = node.add_child(letter[i])
		}
		i++
		if i == n {
			node.is_leaf = true
			node.metadata = metadata
		}
	}
}

func elements(t *Trie, data *[]int) {
	for _, val := range t.children {
		if val.is_leaf {
			*data = append(*data, val.metadata)
		}
		elements(val, data)
	}
}

func (t *Trie) sum(word string) int {
	data := make([]int, 0)
	nt, _ := t.find(word)

	if nt != nil {
		if nt.is_leaf {
			data = append(data, nt.metadata)
		}

		for _, val := range nt.children {
			if val.is_leaf {
				data = append(data, val.metadata)
			}
			elements(val, &data)
		}
	}
	return sum(data)
}

func sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func main() {
	nt := new_trie()
	nt.insert("columnar", 3)
	nt.insert("column", 2)
	fmt.Println("Prefix `col` : ", nt.sum("col"))
	nt_one := new_trie()
	nt_one.insert("bag", 4)
	nt_one.insert("bath", 5)
	fmt.Println("Prefix `ba` :", nt_one.sum("ba"))
}
